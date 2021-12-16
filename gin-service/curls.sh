curl --location --request POST 'http://localhost:8080/files/uploadbinary' \
--header 'Content-Type: application/pdf' \
--data-binary '@/Users/obarra/Downloads/Factura_202jj1-11-03.pdf'

curl --location --request POST 'http://localhost:8080/files/uploadone' \
--form 'file=@"/Users/obarra/Downloads/TuFactura-noviembre.pdf"'

curl --location --request POST 'http://localhost:8080/files/uploadmany' \
--form 'file=@"/Users/obarra/Downloads/TuFactura-noviembre.pdf"' \
--form 'file=@"/Users/obarra/Downloads/d5696ae3-87cb-4e33-a33d-446a50243114.pdf"'