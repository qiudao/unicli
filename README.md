# Description
UNICLI is a Desktop Gui App that mornitor multiple token pair's price at the same time.
The pair price infomation is real time, retrieved from info.uniswap.org
Each pair's info

UNICLI is a Desktop Gui App that displays uniswap's pairs price in realtime. 
It retrieve the market token pair price information from uniswap.org. 
Why need a Desktop Client App?
 some uniswap users want to monitor many token pair price at the same time, manage and 
 group the pairs , look up the pair price history, liquidity history quickly. The Desktop 
 App can support a more high frequency interactive method than the web browser.
 For those users that are accustomed to traditional stock apps, the CLI is very handy.

 The CLI do not support uniswap trade mode, if user want to trade, the CLI support a quick button
 to open the browser page for the token pair on app.uniswap.org.

 In order to reduce the load of uniswap.org, there is a UniServer daemon, which retrieve the pair's
 price from uniswap.org, cache it in redis. The server runs on AWS cloud and Google cloud.

 The UNICLI retrieve data from UniServer, and UniServer retrieve data from uniswap.org.

# Challenge.
 This project include two part: Desktop App and Server daemon.
 The Desktop App use Qt framework and C++ language. The Server daemon use Golang and Redis.
 CLI and Server communicate with protobuffer and GRPC protocol. 
 So the work include:
  - capture data from uniswap.org in realtime.
  - Store the pairs data to Redis.
  - Design the data's scheme for Redis.
  - Design the protobuffer protocol between Server and CLI
  - Design the CLI's UX
  - Qt's adaption for protobuffer.
  - CLI design and implement Price view graph.
  - CLI desgin pairs management and monitor logic.
  - Run on Windows, Linux, Mac os.

<img src="./assets/demo1.gif">
<img src="./assets/demo2.gif">
