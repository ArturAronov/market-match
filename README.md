#### Order Protocol:

Max size: 64 bytes

[1][2][3][4444][5555][6666][77777777][88888888][9999999999999999][10101010101010101010101010101010]

|        | Name             | Offset | Size (Bytes) | Type   | Expected Value                            | Notes / Examples                           |
| ------ | ---------------- | ------ | ------------ | ------ | ----------------------------------------- | ------------------------------------------ |
| **1**  | Transaction Type | 1      | 1            | uint8  | 1 = BUY <br/> 2 = SELL                    | Buy or sell order (0 is invalid)           |
| **2**  | Method           | 2      | 1            | uint8  | 1 = NEW <br/> 2 = MODIFY <br/> 3 = CANCEL | Defines message type (0 is invalid)        |
| **3**  | Order Type       | 3      | 1            | uint8  | 1 = MARKET <br/> 2 = LIMIT                | Market or limit order (0 is invalid)       |
| **4**  | Ticker           | 4      | 4            | string | ASCII, padded with nulls                  | "XYZQ" or "XYZ\0" or "XY\0\0" or "X\0\0\0" |
| **5**  | Quantity         | 8      | 4            | uint32 | 1 – 4,294,967,295                         | Number of shares (must be > 0)             |
| **6**  | Price            | 12     | 4            | uint32 | 1 – 4,294,967,295                         | int in cents, e.g., 12,345 = $123.45       |
| **7**  | Order Date       | 16     | 8            | uint64 | 1 - 18,446,744,073,709,551,615            | Unix epoch (ms)                            |
| **8**  | Good Until       | 24     | 8            | uint64 | 1 - 18,446,744,073,709,551,615            | Unix epoch (ms)                            |
| **9**  | Trader ID        | 32     | 16           | binary | UUID (128-bit)                            | Binary UUID                                |
| **10** | Client Order ID  | 48     | 16           | binary | UUID (128-bit)                            | Binary UUID                                |

#### Return Message Protocol

Max size: 64 bytes

[1][2][3333][4444][5555][6666666666666666][7777777777777777][8888888888888888][9]

|       | Name                 | Offset | Size (Bytes) | Type   | Expected Value                                          | Notes / Examples                           |
| ----- | -------------------- | ------ | ------------ | ------ | ------------------------------------------------------- | ------------------------------------------ |
| **1** | Execution status     | 1      | 1            | uint8  | 1 = PENDING <br/> 2 = PARTIALLY FILLED <br/> 3 = FILLED |                                            |
| **2** | Error                | 2      | 1            | uint8  | 1 = NO ERROR <br/> 2 = WRONG REQUEST                    | Error code (0 is invalid)                  |
| **3** | Quantity Filled      | 3      | 4            | uint32 | 0 – 4,294,967,295                                       | Number of shares filled                    |
| **4** | Avg. Execution Price | 7      | 4            | uint32 | 0 – 4,294,967,295                                       | e.g., 12,345 = $123.45, 0 for non-executed |
| **5** | Execution ID         | 11     | 4            | uint32 | 1 – 4,294,967,295                                       |                                            |
| **6** | Trader ID            | 15     | 16           | binary | UUID (128-bit)                                          | Binary UUID                                |
| **7** | Client Order ID      | 31     | 16           | binary | UUID (128-bit)                                          | Binary UUID                                |
| **8** | Market Order ID      | 47     | 16           | binary | UUID (128-bit)                                          | Binary UUID                                |
| **9** | Padding              | 63     | 1            | byte   | 0                                                       | Reserved for alignment                     |
