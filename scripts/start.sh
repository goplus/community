# run account
cd cmd/gopcomm && cp /config/.community-env .env && go mod tidy && gop run . &
echo "run community"

# run commnuity
cd ../account/cmd && cp /config/.account-env .env && go mod tidy && gop run .
echo "run account"