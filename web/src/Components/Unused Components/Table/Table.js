import React from 'react';
import { Table } from 'reactstrap';

const MemberTable = (props) => {
  return (
    <Table dark hover>
      <thead>
        <tr>
            <th>#</th>
            <th>First Name</th>
            <th>Last Name</th>
            <th>Phone</th>
            <th>Email</th>
            <th>Payment Date</th>
        </tr>
      </thead>
      <tbody>
        <tr>
          <th scope="row">1</th>
          <td>Kevin</td>
          <td>Zeibarth</td>
          <td>985-204-1269</td>
          <td>kevin@amc.com</td>
          <td>8/27/2019</td>
        </tr>
        <tr>
          <th scope="row">2</th>
          <td>Craig</td>
          <td>Canepa</td>
          <td>985-483-5729</td>
          <td>craig@amc.com</td>
          <td>2/11/2019</td>
        </tr>
        <tr>
          <th scope="row">3</th>
          <td>Justin</td>
          <td>Woodring</td>
          <td>225-592-3853</td>
          <td>justin@amc.com</td>
          <td>2/7/2021</td>
        </tr>
      </tbody>
    </Table>
  );
}

export default MemberTable;