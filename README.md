# Load Testing Comparison using Postman

This repository contains the results of a load testing comparison conducted using Postman. The objective of this test was to evaluate the performance of two basic APIs, one written in Golang using Gin and Gorm, and the other in Node.js using Express, when subjected to concurrent GET requests. 

## Test Configuration
- Database: Postgres with 10,000 records.
- Virtual Users: 100.
- Duration: 5 minutes.
- Request Type: GET requests to retrieve all records.
- Results: Available in PDF format in this repository.

## Disclaimer
Please note the following important points:
1. **Non-Scientific Test**: This load testing comparison is not intended to be a scientific benchmark. It was conducted to provide a basic performance overview and should not be considered as a definitive measure of these APIs' capabilities.

2. **Basic APIs**: The APIs tested in this comparison are intentionally kept basic and straightforward. They are not optimized for maximum performance. You may achieve better results by fine-tuning and optimizing your APIs according to your specific requirements.

3. **PDF Results**: The detailed test results are available in PDF format within this repository. Please refer to these files for specific performance data.

## Results
The load testing results can be found in the following PDF files:
- [Results for Golang API](Go-performance-report.pdf)
- [Results for Node.js API](Node-performance-report.pdf)

## Contributions
Contributions and improvements to this load testing comparison are welcome. If you find ways to optimize the APIs or have suggestions for a more comprehensive test, please feel free to open issues or pull requests.

Remember that the goal of this comparison is to provide a basic understanding of performance, and you may be able to build more efficient and high-performing APIs for your specific use case.

## License

This project is licensed under the [MIT License](LICENSE). Feel free to use, modify, and distribute it as needed for your own projects.

Thank you for using and contributing to this Simple CRUD API project! If you have any questions or issues, please don't hesitate to open an issue or reach out to me.
