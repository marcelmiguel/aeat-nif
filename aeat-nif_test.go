package aeatnif

import (
	"log"
	"testing"
)

/*
CERTIFICATE MUST BE REQUESTED TO AEAT
ask password via test or get it via env
*/

const CERTFILE = "cert.crt"
const CERTPWD = ""

//Ok example
func TestServiceOkNIF(t *testing.T) {
	srv := NewVNifV2("https://www1.agenciatributaria.gob.es/wlpl/BURT-JDIT/ws/VNifV2SOAP", CERTFILE, CERTPWD)

	name, _, err := srv.SoapCall("99999999R", "ESPAÑOL ESPAÑOL JUAN")
	if err == nil {
		if *name != "ESPAÑOL ESPAÑOL JUAN" {
			t.Errorf("Data returned is erronious %s", *name)
			log.Println(err)
		}
	} else {
		t.Errorf("Error Accesing SOAP %s : %s", srv.url, err.Error())
		log.Println(err)
	}
}

func TestServicePartialOkNIF(t *testing.T) {
	srv := NewVNifV2("https://www1.agenciatributaria.gob.es/wlpl/BURT-JDIT/ws/VNifV2SOAP", CERTFILE, CERTPWD)

	name, _, err := srv.SoapCall("99999999R", "ESPAÑOL ESPAÑOL") //Requires at lest two words
	if err == nil {
		if *name != "ESPAÑOL ESPAÑOL JUAN" {
			t.Errorf("Data returned is erronious %s", *name)
			log.Println(err)
		}
	} else {
		t.Errorf("Error Accesing SOAP %s : %s", srv.url, err.Error())
		log.Println(err)
	}
}

func TestServiceOkCIF(t *testing.T) {
	srv := NewVNifV2("https://www1.agenciatributaria.gob.es/wlpl/BURT-JDIT/ws/VNifV2SOAP", CERTFILE, CERTPWD)

	name, _, err := srv.SoapCall("B63272603", "")
	if err == nil {
		if *name != "GOOGLE SPAIN, SL" {
			t.Errorf("Data returned is erronious %s", *name)
			log.Println(err)
		}
	} else {
		t.Errorf("Error Accesing SOAP %s : %s", srv.url, err.Error())
		log.Println(err)
	}
}

func TestServiceOkCIF2(t *testing.T) {
	srv := NewVNifV2("https://www1.agenciatributaria.gob.es/wlpl/BURT-JDIT/ws/VNifV2SOAP", CERTFILE, CERTPWD)

	name, taxid, err := srv.SoapCall("C63272603", "")
	if err == nil {
		if (*name != "GOOGLE SPAIN, SL") && (*taxid == "B63272603") {
			t.Errorf("Data returned is erronious %s", *name)
			log.Println(err)
		}
	} else {
		t.Errorf("Error Accesing SOAP %s : %s", srv.url, err.Error())
		log.Println(err)
	}
}

func TestServiceErrorNIF(t *testing.T) {
	srv := NewVNifV2("https://www1.agenciatributaria.gob.es/wlpl/BURT-JDIT/ws/VNifV2SOAP", CERTFILE, CERTPWD)

	name, _, err := srv.SoapCall("99999999S", "ESPAÑOL ESPAÑOL JUAN")
	if err == nil {
		t.Errorf("Error, because it must fail, NIF invalid %s", *name)
		log.Println(err)
	} else {
		log.Println("OK, must fail", err)
	}
}

func TestServiceErrorCIF(t *testing.T) {
	srv := NewVNifV2("https://www1.agenciatributaria.gob.es/wlpl/BURT-JDIT/ws/VNifV2SOAP", CERTFILE, CERTPWD)

	name, _, err := srv.SoapCall("T62880042", "")
	if err == nil {
		t.Errorf("Error, because it must fail, CIF invalid %s", *name)
		log.Println(err)
	} else {
		log.Println("OK, must fail", err)
	}
}

func TestServiceErrorCharSet(t *testing.T) {
	srv := NewVNifV2("https://www1.agenciatributaria.gob.es/wlpl/BURT-JDIT/ws/VNifV2SOAP", CERTFILE, CERTPWD)

	//To correctly test it, chnage xml call to be an incorrect value
	name, _, err := srv.SoapCall("", "")
	if err == nil {
		t.Errorf("Error, because it must fail, NIF invalid %s", *name)
		log.Println(err)
	} else {
		log.Println("OK, must fail", err)
	}
}
