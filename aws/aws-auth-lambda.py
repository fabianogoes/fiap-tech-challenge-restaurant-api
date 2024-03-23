import json

def lambda_handler(event, context):
    print("*********** the event is: *************")
    print(event)
    
    auth = 'Deny'
    if event['authorizationToken'] == 'abc123':
        auth = 'Allow'
    else:
        auth = 'Deny'
        
    authResponse = {}
    authResponse = { "principalId": "abc123", "policyDocument": { "Version": "2012-10-17", "Statement": [{"Action": "execute-api:Invoke", "Resource": ["arn:aws:execute-api:us-east-1:637423526753:6zjg700oqb/*/*/*"], "Effect": auth}] }}
    return authResponse
