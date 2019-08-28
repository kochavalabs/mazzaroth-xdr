%#include "common.x"

namespace mazzaroth
{

  enum BasicType 
  {
    BOOLEAN = 0,
    STRING = 1,
    OPAQUE = 2,
    INT = 3,
    UNSIGNED_INT = 4,
    HYPER = 5,
    UNSIGNED_HYPER = 6,
    FLOAT = 7,
    DOUBLE = 8
  };

  enum ColumnType
  {
    BASIC = 0,
    STRUCT = 1,
    ARRAY = 2
  };

  struct BasicColumn
  {
    string name<40>;

    BasicType typ;
  };

  struct StructColumn
  {
    string name<40>;

    Column columns<40>;
  };


  struct ArrayColumn
  {
    string name<40>;

    boolean fixed;

    unsigned int length;  

    Column column[1];
  };
  
  union Column switch (ColumnType Type)
  {
    case BASIC:
      BasicColumn basic;
    case ARRAY:
      ArrayColumn array;
    case STRUCT:
      StructColumn struc;
  };

  struct Table
  {
    string name<40>;

    Column columns<40>;
  };

  struct Schema
  {
    Table tables<40>;
  };

}
