# Table: locationiq_balance

Account request balance

## Examples

### Get request balance
The Balance API provides a count of request credits left in the user's account for the day. Balance is reset at midnight UTC everyday (00:00 UTC). Bonus is the balance of bonus / promotional request credits in your account.

```sql
select
   balance,
   bonus
from 
   locationiq_balance
```

```
+---------+-------+
| balance | bonus |
+---------+-------+
| 4983    | 0     |
+---------+-------+
```
