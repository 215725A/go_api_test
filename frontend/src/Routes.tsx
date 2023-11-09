import React from 'react';
import { BrowserRouter as Router, Routes, Route, Link } from 'react-router-dom';

import Home from './contents/Home';
import About from './contents/About';
import Test from './contents/Test';
import DB from './contents/DB';

function RoutesComponent() {
    return (
        <Router>
            <div>
                <nav>
                    <ul>
                        <li>
                            <Link to="/">Home</Link>
                        </li>
                        <li>
                            <Link to="/about">About</Link>
                        </li>
                        <li>
                            <Link to="/test">Test</Link>
                        </li>
                        <li>
                            <Link to="/db">DB</Link>
                        </li>
                    </ul>
                </nav>

                <Routes>
                    <Route path="/" element={<Home />} />
                    <Route path="/about" element={<About />} />
                    <Route path="/test" element={<Test />} />
                    <Route path="/db" element={<DB />} />
                </Routes>
            </div>
        </Router>
    );
};

export default RoutesComponent;