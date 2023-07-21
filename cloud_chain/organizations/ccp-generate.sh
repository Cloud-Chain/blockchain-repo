#!/bin/bash

function one_line_pem {
    echo "`awk 'NF {sub(/\\n/, ""); printf "%s\\\\\\\n",$0;}' $1`"
}

function json_ccp {
    local PP=$(one_line_pem $4)
    local CP=$(one_line_pem $5)
    sed -e "s/\${ORG}/$1/" \
        -e "s/\${P0PORT}/$2/" \
        -e "s/\${CAPORT}/$3/" \
        -e "s#\${PEERPEM}#$PP#" \
        -e "s#\${CAPEM}#$CP#" \
        organizations/ccp-template.json
}

function yaml_ccp {
    local PP=$(one_line_pem $4)
    local CP=$(one_line_pem $5)
    sed -e "s/\${ORG}/$1/" \
        -e "s/\${P0PORT}/$2/" \
        -e "s/\${CAPORT}/$3/" \
        -e "s#\${PEERPEM}#$PP#" \
        -e "s#\${CAPEM}#$CP#" \
        organizations/ccp-template.yaml | sed -e $'s/\\\\n/\\\n          /g'
}

ORG=seller
P0PORT=7051
CAPORT=7054
PEERPEM=organizations/peerOrganizations/seller.example.com/tlsca/tlsca.seller.example.com-cert.pem
CAPEM=organizations/peerOrganizations/seller.example.com/ca/ca.seller.example.com-cert.pem


echo "$(json_ccp $ORG $P0PORT $CAPORT $PEERPEM $CAPEM)" > organizations/peerOrganizations/seller.example.com/connection-seller.json

echo "$(yaml_ccp $ORG $P0PORT $CAPORT $PEERPEM $CAPEM)" > organizations/peerOrganizations/seller.example.com/connection-seller.yaml

ORG=buyer
P0PORT=9051
CAPORT=8054
PEERPEM=organizations/peerOrganizations/buyer.example.com/tlsca/tlsca.buyer.example.com-cert.pem
CAPEM=organizations/peerOrganizations/buyer.example.com/ca/ca.buyer.example.com-cert.pem

echo "$(json_ccp $ORG $P0PORT $CAPORT $PEERPEM $CAPEM)" > organizations/peerOrganizations/buyer.example.com/connection-buyer.json
echo "$(yaml_ccp $ORG $P0PORT $CAPORT $PEERPEM $CAPEM)" > organizations/peerOrganizations/buyer.example.com/connection-buyer.yaml


ORG=inspector
P0PORT=11051
CAPORT=11054
PEERPEM=organizations/peerOrganizations/inspector.example.com/tlsca/tlsca.inspector.example.com-cert.pem
CAPEM=organizations/peerOrganizations/inspector.example.com/ca/ca.inspector.example.com-cert.pem

echo "$(json_ccp $ORG $P0PORT $CAPORT $PEERPEM $CAPEM)" > organizations/peerOrganizations/inspector.example.com/connection-inspector.json
echo "$(yaml_ccp $ORG $P0PORT $CAPORT $PEERPEM $CAPEM)" > organizations/peerOrganizations/inspector.example.com/connection-inspector.yaml