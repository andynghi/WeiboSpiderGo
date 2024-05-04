# WeiboSpiderGo

It's a sina weibo (chinese twitter) spider written by golang golly

Weibo crawler that can be run by double-clicking

#### Preparation for use

- chrome driver installation

  You need to install the chrome browser locally and download the corresponding version of chromedriver. For example, if Chrome with version number 78 is installed on this machine, you need to download version number 78 from the link https://chromedriver.chromium.org/downloads, corresponding to the platform's chromedriver.zip

  Place the decompressed chromedriver file in the browser directory of the project

- mongodb installation

  Install mongodb and fill in the address, port, and database name into config.yaml

#### use

After completing the preparations for use in the previous step, you can download the code in the release. In the release page in the upper column, there are compressed packages of the mac version and the windows exe version, which can be downloaded and run directly.

Fill in the accounts that need to be logged in in account/account.txt, one account per line. You can see the example in the account folder of the source code. Double-click login to start batch login (the account must be logged in via an email address that does not require a verification code)

In the account/target.txt file, also one per line, write down the user ID to be crawled, double-click weibo_spider or weibo_spider.exe to start crawling

I hope you won’t fish in the lake. The crawling interval for versions in release is about 10 seconds.

#### Configuration file content

The configuration file is config.yaml in the root directory

Configuration file meaning

```
MONGO_ADDRESS - mongodb address
DB_NAME - mongodb database name
ACCOUNT_FILE - a file that stores the target account id to be crawled
DRIVER_PATH: "/browser/chromedriver"
# Crawl type
SCRAPY_TYPE:
  Info: True
  Follow: False
  Fans: False
  Tweet:
    Main: True
    Comment: False
```







#### Compile

Use after installing dependencies

```
go build login.go
go build weibo_spider.go
```

compile

#### Next step

- [ ] proxy ip added
- [ ] Picture and video download

