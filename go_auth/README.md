# go_auth
Is a small application for personal growth, where the main objective was the JWT and its functioning. Because of this, the errors are not being treated, the ideal would be to treat them according to their context.


## How To Test

1. In your [Postman](https://www.postman.com/home) or [Insomnia](https://insomnia.rest/) restore collection [go_auth.postman_collection.json](https://github.com/piovani/go_api/blob/master/go_auth/go_auth.postman_collection.json)

2. make a request for the "Get Token" route

3. Copy the generated token

4. Add the token to the "Valid Token" route

5. will have two possible results

    <b>Sucess</b>) Token OK

    <b>Fail</b>) Token INVÁLIDO
