package aeatnif

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"text/template"

	"github.com/clbanning/mxj"
)

type VNifV2 struct {
	url          string
	certFile     string
	certPassword string
}

func NewVNifV2(url string, certFile string, certPassword string) *VNifV2 {
	if url == "" {
		url = ""
	}

	return &VNifV2{
		url:          url,
		certFile:     certFile,
		certPassword: certPassword,
	}
}

func generateRequestContent(nif string, name string) string {
	type QueryData struct {
		NIF  string
		NAME string
	}

	const getTemplate = `
	<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:vnif="http://www2.agenciatributaria.gob.es/static_files/common/internet/dep/aplicaciones/es/aeat/burt/jdit/ws/VNifV2Ent.xsd">
	  <soapenv:Header/>
	  <soapenv:Body>
		<vnif:VNifV2Ent>
			<vnif:Contribuyente>
				<vnif:Nif>{{.NIF}}</vnif:Nif>
				<vnif:Nombre>{{.NAME}}</vnif:Nombre>
			</vnif:Contribuyente>
		</vnif:VNifV2Ent>
	  </soapenv:Body>
	</soapenv:Envelope>`

	querydata := QueryData{
		NIF:  nif,
		NAME: name,
	}
	tmpl, err := template.New("requestTemplate").Parse(getTemplate)
	if err != nil {
		panic(err)
	}
	var doc bytes.Buffer
	err = tmpl.Execute(&doc, querydata)
	if err != nil {
		panic(err)
	}

	return doc.String()
}

func convertResults(soapResponse *mxj.Map) (*string, *string, error) {
	successStatus, _ := soapResponse.ValueForPath("Envelope.Body.VNifV2Sal.Contribuyente.Resultado")
	if successStatus != nil {
		success := successStatus.(string) == "IDENTIFICADO"
		if success {
			xmlNombre, _ := soapResponse.ValueForPath("Envelope.Body.VNifV2Sal.Contribuyente.Nombre")
			if xmlNombre == nil {
				return nil, nil, errors.New("Contribuyente.Nombre no encontrado")
			}

			xmlNif, _ := soapResponse.ValueForPath("Envelope.Body.VNifV2Sal.Contribuyente.Nif")
			if xmlNif == nil {
				return nil, nil, errors.New("Contribuyente.Nif no encontrado")
			}

			nombreResult := xmlNombre.(string)
			nifResult := xmlNif.(string)
			return &nombreResult, &nifResult, nil
		} else {
			return nil, nil, errors.New("Identificador fiscal o nombre inválido")
		}
	} else {
		errorMessage, err := soapResponse.ValueForPath("Envelope.Body.Fault.faultstring") //.#text
		if err != nil {
			return nil, nil, errors.New("Certificado erróneo o expirado o servicio de la Agencia Tributaria inaccesible" + err.Error())
		} else {
			return nil, nil, errors.New(errorMessage.(string))
		}
	}
}

func verify(cert *x509.Certificate) error {
	_, err := cert.Verify(x509.VerifyOptions{})
	if err == nil {
		return nil
	}

	switch e := err.(type) {
	case x509.CertificateInvalidError:
		switch e.Reason {
		case x509.Expired:
			return ErrExpired
		default:
			return err
		}
	case x509.UnknownAuthorityError:
		// Apple cert isn't in the cert pool
		// ignoring this error
		return nil
	default:
		return err
	}
}

func (service *VNifV2) SoapCall(nif string, name string) (*string, *string, error) {

	// External certificate
	pemByte, err := ioutil.ReadFile(service.certFile)
	if err != nil {
		return nil, nil, err
	}

	var pemBlocks []*pem.Block
	var v *pem.Block
	var pkey []byte

	for {
		v, pemByte = pem.Decode(pemByte)
		if v == nil {
			break
		}
		/* encrypted certificate
		if v.Type == "RSA PRIVATE KEY" {
			if x509.IsEncryptedPEMBlock(v) {
				pkey, _ = x509.DecryptPEMBlock(v, []byte(service.certPassword))
				pkey = pem.EncodeToMemory(&pem.Block{
					Type:  v.Type,
					Bytes: pkey,
				})
			} else {
				pkey = pem.EncodeToMemory(v)
			}
		} else {
			pemBlocks = append(pemBlocks, v)
		}
		*/
		if v.Type == "PRIVATE KEY" {
			pkey = pem.EncodeToMemory(v)
		} else {
			pemBlocks = append(pemBlocks, v)
		}
	}
	cert, _ := tls.X509KeyPair(pem.EncodeToMemory(pemBlocks[0]), pkey)

	/*err = verify(cert)
	if err != nil {
		return nil, nil, err
	}*/

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				Certificates: []tls.Certificate{cert},
			},
		},
	}

	sRequestContent := generateRequestContent(nif, name)
	requestContent := []byte(sRequestContent)

	req, err := http.NewRequest("POST", service.url, bytes.NewBuffer(requestContent))
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("SOAPAction", `""`)
	req.Header.Add("Content-Type", "text/xml; charset=UTF-8")
	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, nil, errors.New("Fallo al llamar a la agencia tributaria " + resp.Status)
	}
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}
	s := string(contents[:len(contents)])
	fmt.Println(s)
	m, _ := mxj.NewMapXml(contents, true)
	return convertResults(&m)
}

// Certificate errors
var (
	ErrExpired = errors.New("El certificado no es válido")
)
