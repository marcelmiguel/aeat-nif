## AEAT Web Service


# Example responses from agenciatributaria.gob.es

Ok response

``` xml
<?xml version="1.0" encoding="UTF-8"?>
<env:Envelope xmlns:env="http://schemas.xmlsoap.org/soap/envelope/" 
    xmlns:xsd="http://www.w3.org/2001/XMLSchema" 
    xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
    <env:Body>
        <VNifV2Sal:VNifV2Sal xmlns:VNifV2Sal="http://www2.agenciatributaria.gob.es/static_files/common/internet/dep/aplicaciones/es/aeat/burt/jdit/ws/VNifV2Sal.xsd">
            <VNifV2Sal:Contribuyente>
                <VNifV2Sal:Nif>99999999R</VNifV2Sal:Nif>
                <VNifV2Sal:Nombre>ESPAÑOL ESPAÑOL JUAN                     </VNifV2Sal:Nombre>
                <VNifV2Sal:Resultado>IDENTIFICADO</VNifV2Sal:Resultado>
            </VNifV2Sal:Contribuyente>
        </VNifV2Sal:VNifV2Sal>
    </env:Body>
</env:Envelope>
```
Error response

Erroneous call envelope
``` xml
<?xml version="1.0" encoding="UTF-8"?>
<env:Envelope xmlns:env="http://schemas.xmlsoap.org/soap/envelope/" 
    xmlns:xsd="http://www.w3.org/2001/XMLSchema" 
    xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
    <env:Body>
        <VNifV2Sal:VNifV2Sal xmlns:VNifV2Sal="http://www2.agenciatributaria.gob.es/static_files/common/internet/dep/aplicaciones/es/aeat/burt/jdit/ws/VNifV2Sal.xsd">

                <VNifV2Sal:Nif>99999999R</VNifV2Sal:Nif>
                <VNifV2Sal:Nombre>ESPAÑOL ESPAÑOL JUAN                     </VNifV2Sal:Nombre>
                <VNifV2Sal:Resultado>IDENTIFICADO</VNifV2Sal:Resultado>
            </VNifV2Sal:Contribuyente>
        </VNifV2Sal:VNifV2Sal>
    </env:Body>
</env:Envelope>
```

``` xml
<?xml version="1.0" encoding="UTF-8"?>
<env:Envelope xmlns:env="http://schemas.xmlsoap.org/soap/envelope/">
    <env:Body>
        <env:Fault>
            <faultcode>env:Client</faultcode>
            <faultstring>Codigo[1304].El tipo de elemento "vnif:VNifV2Ent" debe terminar con el código de fin correspondiente "&lt;/vnif:VNifV2Ent&gt;". (9,6)</faultstring>
        </env:Fault>
    </env:Body>
</env:Envelope>
```

# Usefull Links

[Read certificate](https://medium.com/@prateeknischal25/using-encrypted-private-keys-with-golang-server-379919955854)

