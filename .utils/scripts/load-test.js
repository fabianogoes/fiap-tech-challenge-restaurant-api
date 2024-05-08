import { check, group } from 'k6';
import http from 'k6/http';

export const options = {
   stages: [
       { duration: '0.2m', target: 3 }, // simulate ramp-up of traffic from 1 to 3 virtual users over 0.5 minutes.
       { duration: '0.2m', target: 4}, // stay at 4 virtual users for 0.5 minutes
       { duration: '0.2m', target: 0 }, // ramp-down to 0 users
     ],
};

export default function () {
   group('API uptime heath check', () => {
       const response = http.get('http://localhost:8080/health');
       check(response, {
           "status code should be 200": res => res.status === 200,
       });
   });

   group('API uptime GET attendants', () => {
       const response = http.get('http://localhost:8080/attendants');
       check(response, {
           "status code should be 200": res => res.status === 200,
       });
   });

   group('API uptime GET customers', () => {
       const response = http.get('http://localhost:8080/customers');
       check(response, {
           "status code should be 200": res => res.status === 200,
       });
    }); 
    
    group('API uptime GET products', () => {
        const response = http.get('http://localhost:8080/products');
        check(response, {
            "status code should be 200": res => res.status === 200,
        });
    }); 
    
    let orderID;
     group('Create a Order', () => {
        const url = 'http://localhost:8080/orders';
        const payload = JSON.stringify({"customerCPF": "15204180001", "attendantID": 1});
        const params = {
            headers: {
                'Content-Type': 'application/json',
            },
        };
        const response = http.post(url, payload, params);
        orderID = response.json().id;
        check(response, {
            "status code should be 201": res => res.status === 201,
        });
     })

     group('Add Item to Order', () => {
        const url = `http://localhost:8080/orders/${orderID}/item`;
        const payload = JSON.stringify({"productID": 1, "quantity": 2});
        const params = {
            headers: {
                'Content-Type': 'application/json',
            },
        };
        const response = http.post(url, payload, params);
        check(response, {
            "status code should be 201": res => res.status === 201,
        });
     })

     group('Confirmation a Order', () => {
        const response = http.put(`http://localhost:8080/orders/${orderID}/confirmation`);
        check(response, {
            "status code should be 200": res => res.status === 200,
        });
     })

     group('Payment a Order', () => {
        const url = `http://localhost:8080/orders/${orderID}/payment`;
        const payload = JSON.stringify({"paymentMethod": "DEBIT_CARD"});
        const params = {
            headers: {
                'Content-Type': 'application/json',
            },
        };
        const response = http.put(url, payload, params);
        check(response, {
            "status code should be 200": res => res.status === 200,
        });
     })

     group('Send a Order to Preparation', () => {
        const response = http.put(`http://localhost:8080/orders/${orderID}/in-preparation`);
        check(response, {
            "status code should be 200": res => res.status === 200,
        });
     })

     group('Send Order to Ready for Delivery', () => {
        const response = http.put(`http://localhost:8080/orders/${orderID}/ready-for-delivery`);
        check(response, {
            "status code should be 200": res => res.status === 200,
        });
     })
     
     group('Sent Order for Delivery', () => {
        const response = http.put(`http://localhost:8080/orders/${orderID}/sent-for-delivery`);
        check(response, {
            "status code should be 200": res => res.status === 200,
        });
     })

     group('Order mark as Delivered', () => {
        const response = http.put(`http://localhost:8080/orders/${orderID}/delivered`);
        check(response, {
            "status code should be 200": res => res.status === 200,
        });
     })

};

