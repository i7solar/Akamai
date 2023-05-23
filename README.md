# Akamai 1.75 Generator 🤖

## Introduction
Akamai 1.75 is a robust and high-performance tool, developed with Go, dedicated to generating Akamai cookies. Akamai serves as a comprehensive and specialized bot protection system that fortifies an array of websites against various threats. Throughout its tenure, this API has effectively catered to more than five million requests, highlighting its reliability and proficiency. This generator will only work on Akamai 1.7X (legacy) websites.

## Acknowledgements
This project signifies my first venture into the dynamic realm of Software Engineering. Thank you to Levi for helping with business-logic/configuring AWS, Kalek for helping with devices, Kiwi for Go & Akamai questions, and others for their invaluable contributions and continuous support.

## Project Origin
This endeavor was kick-started two years ago, with a primary objective to devise stable Akamai cookies for SolarSystems Software. This software was an e-commerce bot engineered to streamline and automate the buying process, thereby providing a robust solution for retail arbitrage.

## Akamai Bot Manager
Akamai Bot Manager is an integral component of the Akamai Intelligent Edge Platform, offering a comprehensive solution for managing bots' impact on websites and mobile applications. Equipped with an array of tools, the Bot Manager efficiently distinguishes between beneficial bots, malicious bots, and human users, thereby facilitating optimal website performance and user experience.

## Technical Specifications
* **Script Porting:** The core JavaScript script was translated to Go, augmenting the application's performance and scalability.
* **Simulated User Interactions:** Our tool employs simulated timings and mouse movements using collected data for enhanced realism. However, it's crucial to note that the simulated mouse movements no longer bypass the updated Akamai protection.
* **Architecture and Hosting:** The application, built on Go Fiber, is architected to operate seamlessly on AWS EC2/API Gateway.

## Deployment
Please note, you must adapt the HTTP client to meet your specific requirements before proceeding. Post adaptation, you may execute the Go Fiber application using `main.go`.

* Ensure that Go is installed on your machine prior to running the application.

```bash
# Clone the project source code
git clone https://github.com/i7solar/Akamai.git

# Navigate to the project directory
cd Akamai

# Execute the application
go run main.go
```

## Contributing
We encourage developers to contribute and help us improve this project's effectiveness. Feel free to open issues, suggest changes, and make the project more beneficial for everyone.
