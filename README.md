# ForgeRock Access Manager Client

## curl

```
ACCESSTOKEN=`curl -X POST \
    -H "X-OpenAM-Username: ${FRAM_USERNAME}" \
    -H "X-OpenAM-Password: ${FRAM_PASSWORD}" \
    -H "Accept-API-Version: resource=2.1" \
    https://fram.darkedges.com/openam/json/realms/root/authenticate | jq -r .tokenId`
```

```
curl -X GET \
    "http://fram.darkedges.com/openam/json/realm-config/services/baseurl" \
        -H "Cookie: iPlanetDirectoryPro=${ACCESSTOKEN}" | jq .
```

```
curl -X PUT \
    "http://fram.darkedges.com/openam/json/realm-config/services/baseurl" \
        -H "Content-Type: application/json" \
        -H "X-Requested-With: SwaggerUI" \
        -H "Cookie: iPlanetDirectoryPro=${ACCESSTOKEN}" \
        -d "{ \"source\": \"FIXED_VALUE\", \"fixedValue\": \"https://fram.darkedges.com\", \"extensionClassName\": \"string\", \"contextPath\": \"/openam\"}" | jq .
```