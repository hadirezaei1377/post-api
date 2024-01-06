CREATE TABLE History (
    Id SERIAL PRIMARY KEY,
    HistoryId int REFERENCES History(Id),
    Date timestamptz,
    Open decimal,
    High decimal,
    Low decimal,
    Close decimal
);

id is unique for each instrument
History(Id) (int) is related to instrument
Date (timestamptz) is about date and time of instrument
Open (decimal) the price for beginning
High (decimal) the most expensive value for instrument
Low (decimal) the cheapest value for instrument
Close (decimal) the much of instrument when it was finishing