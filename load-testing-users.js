import http from 'k6/http';
import { uuidv4 } from 'https://jslib.k6.io/k6-utils/1.4.0/index.js';
import { check } from 'k6';

export const options = {
  vus: 50,
  duration: '20s',
};

export default function () {
  const randomUUID = uuidv4();
  const payload = JSON.stringify({
    name: 'Luis',
    email: `${String(randomUUID)}@gmail.com`,
    telephone: '3123202',
    id: String(randomUUID),
  });
  const res = http.post('http://localhost:9091/api/v1/users', payload);
  check(res, {
    "is OK": (r)=>{
      return r.status == 201
    }
  })
}
