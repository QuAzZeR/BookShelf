import React from 'react';
import { Table } from 'antd';
const TableBooks = (data, columns) => (
  <Table
    dataSource={data}
    columns={columns}
  >

  </Table>
)