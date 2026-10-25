#!/usr/bin/env bash

set -e

BASE_URL="http://localhost:8080"

SEPARATOR="================================="

echo "$SEPARATOR"
echo "1. HEALTH CHECK"
echo "$SEPARATOR"

curl -s \
${BASE_URL}/healthz

echo
echo

##############################################################################
# LOGIN ADMIN
##############################################################################

echo "$SEPARATOR"
echo "2. LOGIN ADMIN"
echo "$SEPARATOR"

LOGIN_RESPONSE=$(curl -s \
-X POST ${BASE_URL}/auth/login \
-H "Content-Type: application/json" \
-d '{
  "email":"admin@company.com",
  "password":"admin123"
}')

echo "$LOGIN_RESPONSE"

TOKEN=$(echo "$LOGIN_RESPONSE" \
| sed -n 's/.*"token":"\([^"]*\)".*/\1/p')

if [ -z "$TOKEN" ]; then
  echo "ERROR: ADMIN TOKEN NOT FOUND"
  exit 1
fi

echo
echo "ADMIN TOKEN OK"
echo

##############################################################################
# CREATE USER
##############################################################################

TIMESTAMP=$(date +%s)

USER_EMAIL="john${TIMESTAMP}@company.com"

echo "$SEPARATOR"
echo "3. CREATE USER"
echo "$SEPARATOR"

CREATE_USER_RESPONSE=$(curl -s \
-X POST ${BASE_URL}/api/v1/users \
-H "Authorization: Bearer ${TOKEN}" \
-H "Content-Type: application/json" \
-d "{
  \"email\":\"${USER_EMAIL}\",
  \"password\":\"Password123\",
  \"firstname\":\"John\",
  \"lastname\":\"Doe\",
  \"role\":\"manager\"
}")

echo "$CREATE_USER_RESPONSE"

echo
echo

##############################################################################
# LIST USERS
##############################################################################

echo "$SEPARATOR"
echo "4. LIST USERS"
echo "$SEPARATOR"

USERS_RESPONSE=$(curl -s \
${BASE_URL}/api/v1/users \
-H "Authorization: Bearer ${TOKEN}")

echo "$USERS_RESPONSE"

echo
echo

##############################################################################
# LOGIN JOHN
##############################################################################

echo "$SEPARATOR"
echo "5. LOGIN JOHN"
echo "$SEPARATOR"

JOHN_LOGIN=$(curl -s \
-X POST ${BASE_URL}/auth/login \
-H "Content-Type: application/json" \
-d "{
  \"email\":\"${USER_EMAIL}\",
  \"password\":\"Password123\"
}")

echo "$JOHN_LOGIN"

JOHN_TOKEN=$(echo "$JOHN_LOGIN" \
| sed -n 's/.*"token":"\([^"]*\)".*/\1/p')

if [ -z "$JOHN_TOKEN" ]; then
  echo "ERROR: JOHN TOKEN NOT FOUND"
  exit 1
fi

echo
echo

##############################################################################
# USER ID
##############################################################################

USER_ID=$(echo "$USERS_RESPONSE" \
| grep -o '"id":"[^"]*"' \
| head -1 \
| cut -d':' -f2 \
| tr -d '"')

echo "$SEPARATOR"
echo "6. GET USER"
echo "$SEPARATOR"

curl -s \
${BASE_URL}/api/v1/users/${USER_ID} \
-H "Authorization: Bearer ${TOKEN}"

echo
echo

##############################################################################
# UPDATE USER
##############################################################################

echo "$SEPARATOR"
echo "7. UPDATE USER"
echo "$SEPARATOR"

curl -s \
-X PUT \
${BASE_URL}/api/v1/users/${USER_ID} \
-H "Authorization: Bearer ${TOKEN}" \
-H "Content-Type: application/json" \
-d '{
  "role":"manager",
  "active":true
}'

echo
echo

##############################################################################
# CREATE CUSTOMER
##############################################################################

CUSTOMER_EMAIL="contact${TIMESTAMP}@acme.com"

echo "$SEPARATOR"
echo "8. CREATE CUSTOMER"
echo "$SEPARATOR"

curl -s \
-X POST ${BASE_URL}/api/v1/customers \
-H "Authorization: Bearer ${JOHN_TOKEN}" \
-H "Content-Type: application/json" \
-d "{
  \"name\":\"Acme Corp\",
  \"email\":\"${CUSTOMER_EMAIL}\",
  \"phone\":\"+22670000000\"
}"

echo
echo

##############################################################################
# LIST CUSTOMERS
##############################################################################

echo "$SEPARATOR"
echo "9. LIST CUSTOMERS"
echo "$SEPARATOR"

CUSTOMERS=$(curl -s \
${BASE_URL}/api/v1/customers \
-H "Authorization: Bearer ${JOHN_TOKEN}")

echo "$CUSTOMERS"

echo
echo

##############################################################################
# CUSTOMER ID
##############################################################################

CUSTOMER_ID=$(echo "$CUSTOMERS" \
| grep -o '"id":"[^"]*"' \
| head -1 \
| cut -d':' -f2 \
| tr -d '"')

echo "CUSTOMER_ID=${CUSTOMER_ID}"
echo

##############################################################################
# GET CUSTOMER
##############################################################################

echo "$SEPARATOR"
echo "10. GET CUSTOMER"
echo "$SEPARATOR"

curl -s \
${BASE_URL}/api/v1/customers/${CUSTOMER_ID} \
-H "Authorization: Bearer ${JOHN_TOKEN}"

echo
echo

##############################################################################
# UPDATE CUSTOMER
##############################################################################

echo "$SEPARATOR"
echo "11. UPDATE CUSTOMER"
echo "$SEPARATOR"

curl -s \
-X PUT \
${BASE_URL}/api/v1/customers/${CUSTOMER_ID} \
-H "Authorization: Bearer ${JOHN_TOKEN}" \
-H "Content-Type: application/json" \
-d '{
  "name":"Acme Corporation",
  "status":"active"
}'

echo
echo

##############################################################################
# CREATE CONTACT
##############################################################################

echo "$SEPARATOR"
echo "12. CREATE CONTACT"
echo "$SEPARATOR"

curl -s \
-X POST ${BASE_URL}/api/v1/contacts \
-H "Authorization: Bearer ${JOHN_TOKEN}" \
-H "Content-Type: application/json" \
-d "{
  \"customer_id\":\"${CUSTOMER_ID}\",
  \"firstname\":\"Jean\",
  \"lastname\":\"Dupont\",
  \"email\":\"jean${TIMESTAMP}@acme.com\",
  \"phone\":\"+22670222222\"
}"

echo
echo

##############################################################################
# LIST CONTACTS
##############################################################################

echo "$SEPARATOR"
echo "13. LIST CONTACTS"
echo "$SEPARATOR"

curl -s \
${BASE_URL}/api/v1/contacts \
-H "Authorization: Bearer ${JOHN_TOKEN}"

echo
echo

##############################################################################
# AUDIT LOGS
##############################################################################

echo "$SEPARATOR"
echo "14. AUDIT LOGS"
echo "$SEPARATOR"

curl -s \
${BASE_URL}/api/v1/audit \
-H "Authorization: Bearer ${TOKEN}"

echo
echo

##############################################################################
# INVALID TOKEN
##############################################################################

echo "$SEPARATOR"
echo "15. INVALID TOKEN"
echo "$SEPARATOR"

curl -s \
${BASE_URL}/api/v1/users \
-H "Authorization: Bearer invalid-token"

echo
echo

##############################################################################
# NO TOKEN
##############################################################################

echo "$SEPARATOR"
echo "16. NO TOKEN"
echo "$SEPARATOR"

curl -s \
${BASE_URL}/api/v1/users

echo
echo

##############################################################################
# RBAC TEST
##############################################################################

echo "$SEPARATOR"
echo "17. RBAC"
echo "$SEPARATOR"

curl -s \
-X DELETE \
${BASE_URL}/api/v1/users/${USER_ID} \
-H "Authorization: Bearer ${JOHN_TOKEN}"

echo
echo

##############################################################################
# DONE
##############################################################################

echo "$SEPARATOR"
echo "ALL TESTS FINISHED"
echo "$SEPARATOR"
