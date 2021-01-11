# StockScans
StockScans uses the Alpaca API to find opportunities in the stock market through various trading strategies

# Supported Scan Types
  1. Inside Bars
  Coming soon: Oversold Bounces (RSI)
  
# Supported Time Frames
  Weekly, Daily, Hourly, 15-minute, 5-minute
  
# How to use
  1. Launch application for the first time 
  2. A config.yml should be generated and the application will close. Edit the config.yml file with your https://alpaca.markets API key-id and key-secret.
  3. Launch application again. Select your options.
  4. The tickers that match your requirements will be outputted to the console.
  
# Known Issues
  1. Handling partial bars on Inside Bars mode (such as scanning for weekly inside bars on a Tuesday) is not fully supported and may cause some issues. 
     (However, the scanner will likely still return tight/compressed ranges, except they might not fit the exact definition of an IB)
  
