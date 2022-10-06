# expentracker
One stop solution for all your financial planning


# brainstorming
write a do while loop to check the token is present or not from redis cache
another gin endpoint will get and save the code to redis cache 

login flow - 
call dashboard from js with auth code in header
validate authcode from header
if invalid return with error and display below button a msg "Login failed try again"
if valid then get all the user data from DB
(if user does not exist send error saying user does not exist with error code and display complete your profile data )(hidden)
if user does not exist send the UserData model with profile info and 0 acc name and 0 cards names
if user exist then send the whole UserData model data like user data and account name and card names

design a [refresh] button you will send request to fetchexpenditure endpoint there we will get the google oauth 2 request url then we will redirect to the consent page and get the access code from the callback. now from access code we will exchange an access token this access token will be saved in redis for convenience as we will save its key as email id like emailid:access_token otherwise we would have stored access token as id:token pair