CREATE TABLE IF NOT EXISTS delivery (
    Id      int PRIMARY KEY,
    Name    text,
    Phone   text,
    Zip     text,
    City    text,
    Address text,
    Region  text,
    Email   text
);

CREATE TABLE IF NOT EXISTS payment (
    transaction  text PRIMARY KEY,
    RequestId    text,
    Currency     text,
    Provider     text,
    Amount       int,
    PaymentDt    int,
    Bank         text,
    DeliveryCost int,
    GoodsTotal   int,
    CustomFee    int
);

CREATE TABLE IF NOT EXISTS orders (
    OrderUid          text PRIMARY KEY,
    TrackNumber       text,
    Entry             text,
    DeliveryId        int REFERENCES delivery,
    PaymentId         text REFERENCES payment,
    Locale            text,
    InternalSignature text,
    CustomerId        text,
    DeliveryService   text,
    Shardkey          text,
    SmId              int,
    DateCreated       text,
    OofShard          text
);

CREATE TABLE IF NOT EXISTS item (
    rid         text    PRIMARY KEY,
    ChrtId      int,
    TrackNumber text,
    Price       int,
    Name        text,
    Sale        int,
    Size        text,
    TotalPrice  int,
    NmId        int,
    Brand       text,
    Status      int,
    OrderUid    text REFERENCES orders
);
