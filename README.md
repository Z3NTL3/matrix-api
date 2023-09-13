# Distance Calculator API
Calculate the distance between two variating locations in KM, also provides an duration overview in minutes.

**This is an wrapper, when the original website goes down because of exceeded quota, then download and serve this server and drop your token into .env**
> openrouteservice.org

### Notice
This is something i have programmed for a few fellows on my school. So that it helps in their current project.


> So for the reason i mentioned above, i have not implemented very advanced security stuff, only honeypot reverse blocking at reverse proxy level.

## Backend
- Programmed in `Go Gin`

# API REQUESTS
**Example:**
```
[ROUTE]
GET https://distance.pix4.dev/api/calculate/distance/:origin/:dest

[EXAMPLE REQUEST]
GET https://distance.pix4.dev/api/calculate/distance/Doetinchem, Nederland/Tivolilaan 40, 6824 BW Arnhem, Nederland

[API RESPONSE EXAMPLE]
{"Success":true,"Data":{"Duration":"33.29 min","Distance":"33.96 KM"}}
```