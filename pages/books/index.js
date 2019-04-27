import React from 'react';
import {Table} from 'antd';
import './index.less';

const columns = [{
  title: 'Name',
  dataIndex: 'name',
  key: 'name',
}, {
  title: 'Age',
  dataIndex: 'age',
  key: 'age',
}, {
  title: 'Address',
  dataIndex: 'address',
  key: 'address',
}];

const dataSource = [{
  key: '1',
  name: 'Mike',
  age: 32,
  address: '10 Downing Street',
}, {
  key: '2',
  name: 'John',
  age: 42,
  address: '10 Downing Street',
}];

const BookPages = () => (
  <div className="BookPages__container">
    <Table
      dataSource={dataSource}
      columns={columns}
    />

  </div>
);

export default BookPages;