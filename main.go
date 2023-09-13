/*
Programmed by Z3NTL3
Also known as Efdal
*/

package main

import (
	"fmt"
	"log"
	"strconv"
	"time"
	"z3ntl3/distance-api/bot"
	"z3ntl3/distance-api/globals"
	"z3ntl3/distance-api/mw"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Cannot read '.env'")
	}

	gin.SetMode(gin.ReleaseMode)
	app := gin.Default()
	{
		//app.RemoteIPHeaders = []string{"cf-connecting-ip"}
		app.UseH2C = true
		app.RedirectTrailingSlash = true

		app.Use(mw.Credits)

		{
			app.NoRoute(func(ctx *gin.Context) {
				ctx.JSON(404, &globals.API_Resp{
					Success: false,
					Data: struct{ Message string }{
						Message: fmt.Sprintf("[%s] Route not found", ctx.Request.URL),
					},
				})
			})

			app.NoMethod(func(ctx *gin.Context) {
				ctx.JSON(404, &globals.API_Resp{
					Success: false,
					Data: struct{ Message string }{
						Message: fmt.Sprintf("[%s] Method not allowed", ctx.Request.Method),
					},
				})
			})

			// had geen zin om rate limiter heb honeypot block met cf proxies so idc.
			api := app.Group("/api")
			{
				api.GET("/calculate/distance/:origin/:dest", func(ctx *gin.Context) {
					origin, ok := ctx.Params.Get("origin")
					if !ok {
						ctx.JSON(400, globals.API_Resp{
							Success: false,
							Data: struct{ Message string }{
								Message: "parameter origin is not even set dude?!!!!........",
							},
						})
						return
					}
					dest, ok := ctx.Params.Get("dest")
					if !ok {
						ctx.JSON(400, globals.API_Resp{
							Success: false,
							Data: struct{ Message string }{
								Message: "parameter dest is not even set dude?!!!!........",
							},
						})
						return
					}

					token := ctx.Query("token")

					data, err := bot.RunBot(origin, dest, token)
					if err != nil {
						//fmt.Println("hiero", err)
						ctx.JSON(500, globals.API_Resp{
							Success: false,
							Data: struct{ Message string }{
								Message: err.Error(),
							},
						})
						return
					}

					bigBoiii := struct {
						Duration string
						Distance string
					}{} // nil for now

					{
						// big boii stuff xxxRRR
						fix := strconv.FormatFloat(data.Duration, 'f', -1, 64)
						inMinutes, err := time.ParseDuration(fix + "s")
						if err != nil {
							ctx.JSON(500, globals.API_Resp{
								Success: false,
								Data: struct{ Message string }{
									Message: err.Error(),
								},
							})
							return
						}

						(&bigBoiii).Duration = fmt.Sprintf("%.2f min", inMinutes.Minutes())
						(&bigBoiii).Distance = strconv.FormatFloat(data.Distance, 'f', -1, 64) + " KM"
					}

					ctx.JSON(200, globals.API_Resp{
						Success: true,
						Data:    bigBoiii,
					})
				})
			}
		}
	}

	app.Run(":2000")
}
