import React from 'react';
import {
  BrowserRouter as Router,
  Routes,
  Route,
} from "react-router-dom";

import Navbar from 'react-bootstrap/Navbar';
import Nav from 'react-bootstrap/Nav';
import Form from 'react-bootstrap/Form';
import FormControl from 'react-bootstrap/FormControl';
import Button from 'react-bootstrap/Button';

import './App.css';

import ProductList from './ProductList.js';

import 'bootstrap/dist/css/bootstrap.min.css';
import { useParams } from "react-router-dom";
import UploadImage from './UploadImage.js';

function UploadImageWrapper(props) {
  const params = useParams();
  return <UploadImage {...props} params={params} />;
}

function App() {
  return (
    <Router>
      <div className="App">
        <Navbar bg="light" expand="lg">
          <Navbar.Brand href="/">Coffee Shop</Navbar.Brand>
          <Navbar.Toggle aria-controls="basic-navbar-nav" />
          <Navbar.Collapse id="basic-navbar-nav">
          </Navbar.Collapse>
        </Navbar>

        <Routes>
          <Route path="/upload-image/:id" element={<UploadImageWrapper />} />
          <Route path="/" element={<ProductList />} />
          <Route path="*" element={<ProductList />} />
        </Routes>
      </div>
    </Router>
  );
}

export default App;
