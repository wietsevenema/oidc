# oidc
A OpenID Connect (OIDC) token parser and pretty printer in one file - 
code you can read. 

Friends don't let friends copy & paste tokens on untrusted websites to 
learn what is inside.  

Summarizing this little package:
 - It decodes the OIDC token and pretty prints header and claims. 
 - There are no (zero) dependencies.
 - The package is just one main.go with code you can (and should!) read.
 - Does not verify signature - use only for debugging. 
 
# Why?
This package is useful if you are trying to debug or understand OIDC tokens.

On Google Cloud, more and more APIs support OpenID Connect 
(OIDC) as one of the primary authentication methods, for both user accounts
and service accounts.

The token is your means of authentication, and you don't want to share
it with strangers. This is exactly why I made this utility, and why it may
also be useful to you.   

# What is a OIDC token?
You use an OIDC token to prove your identity to a server 
you want to access, sending the token using the `Authorization: Bearer` HTTP header. 

The OIDC token is an encoded (not encrypted) object with a cryptographic signature. 
It has three parts: the header, the claims (body with values), 
and a signature.

The *fields* in the tokens are specified by the OIDC standard, and 
the *format* of the token is a JSON Web Token (JWT).

The token contains the destination URL (audience), the identity of the caller 
(email), and an expiration date. You can’t get tokens that are valid for longer
than one hour (a comforting thought). You get a token, and pass it to the API 
you want to call.

# Installation
`go get github.com/wietsevenema/oidc`

# Example Usage
In the next snippet, I am printing my personal token using `gcloud auth print-identity-token`. 
When you give this token to someone else, they can access services 
on your behalf for an hour. This is why you should never paste the raw token on a 
random website or trust code from a stranger to handle it for you. Please read the 
code in main.go to verify I am not doing anything suspicious.  
```
$: gcloud auth print-identity-token | oidc
{
  "Header": {
    "alg": "RS256",
    "kid": "cb404383844b46312769bb929ecec57d1ad8e3bb",
    "typ": "JWT"
  },
  "Claims": {
    "at_hash": "MfiwIy0I-O72D6JBiYiCLw",
    "aud": "[SNIP]",
    "azp": "[SNIP]",
    "email": "[SNIP]",
    "email_verified": true,
    "exp": 1584025433,
    "iat": 1584021833,
    "iss": "https://accounts.google.com",
    "sub": "105805474220960633511"
  }
}
```
