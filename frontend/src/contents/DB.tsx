import React, { useEffect, useState } from 'react';

const DB = () => {
    const [data, setData] = useState(null);

    useEffect(() => {
        fetch("http://localhost:8080/api/db")
            .then((response) => response.json())
            .then((data) => setData(data))
            .catch((error) => console.error('Error fetching data:' + error));
    }, []);

    return (
        <div>
            <h2>Data from Backend API</h2>
            {data && <pre>{JSON.stringify(data, null, 2)}</pre>}
        </div>
    );
};

export default DB;