import Container from 'react-bootstrap/Container';
import Nav from 'react-bootstrap/Nav';
import Navbar from 'react-bootstrap/Navbar';
import NavDropdown from 'react-bootstrap/NavDropdown';
import Imagex from '../assets/hacker.png';
import './navbar.css';

export default function Navbarr() {
  return (
    <nav className="navbar">
      <div className="navbar-brand-container">
        <img src={Imagex} alt="image" />
        <div className="navbar-brand">Task Manager</div>
      </div>
      <ul className="navbar-links">
        <li className="navbar-item"><a href="#home" className="navbar-link">Home</a></li>
        <li className="navbar-item"><a href="#about" className="navbar-link">Task</a></li>
      </ul>
    </nav>
  );
}
