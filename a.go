return http.HandleFunc (func (w http.Responsewriter, r *http.Request){
    if r. Header["Token"] != nil {
        jwt.Parse(r.Header["Token"][0], func(token *jwt. Token) (interface(), error){
            if , ok :- token. Method. (*jwt.SigningMethodHMAC); lok{
                 return nil, fmt. Errorf (("Invalid signing method"))
            aud := "billing.jwtgo.io"
            checkAudience := token.Claims. (jwt. Mapclaims). VerifyAudience(aud, false)
            if IcheckAudience (
                 return nil, fmt. Errorf(("invalid aud"))
            iss := "jwtgo.io"
            checkIss := token.Claims. (jwt.MapClaims).VerifyIssuer(iss, false)
            if !checkIss (
                 return nil, fmt. Errorf (("invalid iss"))
            return Mysigningkey, nil
        })

	}