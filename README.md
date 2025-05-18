# salesAnalytics

This project processes a large CSV file containing historical sales data, potentially consisting of millions of rows.

# Technology Stack

- Go (Golang): Backend APIs and CSV processing
- GORM: ORM for database interactions
- PostgreSQL: Database to store sales records
- JSON API: For frontend/backend communication

## External Package

Toml 
GORM 
MUX

## Install Dependencies

go mod tidy

## Run

go run main.go

## LOG file 
The application creates logs while running to help track the execution flow and debug issues.
Log Location: log/logfile19042025.12.20.03.610924298.txt

## Api

Url : http://localhost:23434/based_onRevenue

Input value :
Body 
{
   "fromDate": "2023-12-15",
   "endDate": "2024-05-18"
}
Header :
Indicator : "Date_range"

On Error :
{
    "status": "E",
    "statusCode": "GRD01",
    "msg": "EOF"
}

ON Success :
Response 1 :
{
    "status": "S",
    "total_Revenue": "4448.5675",
    "totProdRevenue": null,
    "totalcatRevenue": null,
    "totalregionRevenue": null,
    "topProduct": null,
    "topcategory": null,
    "topRegion": null
}

Response : 2
{
    "status": "S",
    "total_Revenue": "",
    "totProdRevenue": null,
    "totalcatRevenue": [
        {
            "category": "Clothing",
            "total_revenue": 143.976
        },
        {
            "category": "Electronics",
            "total_revenue": 4064.5915
        },
        {
            "category": "Shoes",
            "total_revenue": 180
        }
    ],
    "totalregionRevenue": null,
    "topProduct": null,
    "topcategory": null,
    "topRegion": null
}

Reponse 3 :
{
    "status": "S",
    "total_Revenue": "",
    "totProdRevenue": null,
    "totalcatRevenue": null,
    "totalregionRevenue": [
        {
            "region": "",
            "total_revenue": 2612.076
        },
        {
            "region": "",
            "total_revenue": 1299
        },
        {
            "region": "",
            "total_revenue": 297.4915
        },
        {
            "region": "",
            "total_revenue": 180
        }
    ],
    "topProduct": null,
    "topcategory": null,
    "topRegion": null
}

Url : http://localhost:23434/Get_Productsdetails

ON Error :
{
    "status": "E",
    "statusCode": "GPD02",
    "msg": "Error 1054 (42S22): Unknown column 'products.product_name' in 'SELECT'"
}

ON Success :
{
    "status": "S",
    "total_Revenue": "",
    "totProdRevenue": null,
    "totalcatRevenue": null,
    "totalregionRevenue": null,
    "topProduct": [
        {
            "product_name": "",
            "total_quantity": 3
        },
        {
            "product_name": "",
            "total_quantity": 3
        },
        {
            "product_name": "",
            "total_quantity": 1
        },
        {
            "product_name": "",
            "total_quantity": 1
        }
    ],
    "topcategory": null,
    "topRegion": null
}

{
    "status": "S",
    "total_Revenue": "",
    "totProdRevenue": null,
    "totalcatRevenue": null,
    "totalregionRevenue": null,
    "topProduct": null,
    "topcategory": [
        {
            "category": "Electronics",
            "total_quantity": 4
        }
    ],
    "topRegion": null
}