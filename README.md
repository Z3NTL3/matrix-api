# Distance Matrix Wrapper API
Calculate the distance between two variating locations in KM, also provides a duration overview in minutes.

**This is a wrapper, by default this API uses it's own API key. When the quota exceeds the limit, you may provide your own token by adding querystring 'token' to our API. Below i've mentioned the website from where to get your API key**
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
GET https://distance.pix4.dev/api/calculate/distance/:origin/:dest/<optional>?token=your_token

[EXAMPLE REQUEST]
GET https://distance.pix4.dev/api/calculate/distance/Amsterdam, Nederland/Tivolilaan 40, 6824 BW Arnhem, Nederland

```
```json
[API RESPONSE EXAMPLE]
{"Success":true,"Data":{"Duration":"81.95 min","Distance":"101.01 KM"}}
```
