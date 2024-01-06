Query to report the last transaction of each image:

SELECT
    Image.Name AS Symbol,
    History.Date AS TransactionDate,
    Trade.Open,
    Trade.High,
    Trade.Low,
    Trade.Close
FROM
    Image
JOIN LATERAL (
    SELECT *
    FROM History
    WHERE ImageId = Image.Id
    ORDER BY Date DESC
    LIMIT 1
) AS Trade ON true;



This query will show us the information of the last trade of each symbol based on the trade date.
 In the sample data, a history of transactions is inserted exponentially by date.

 Simplification:
 SELECT Instrument.Name AS Symbol : This section is for selecting the desired columns from the tables
 Instrument.Name AS Symbol : displays the name of the symbol as a symbol
 Trade.DateEn AS TransactionDate : displays the transaction date as TransactionDate
 FROM Instrument : This section shows the desired table for information extraction (Instrument)
ImageId = Image.Id : This section specifies that the symbol in the Trade table matches the Id of the symbol in the Instrument table
