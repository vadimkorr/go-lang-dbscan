//    Y
//    /\
//    |
// 8  |
// 7  |   p0(2,6)                          p6(12,6)
// 6  |       *                             *
// 5  |   p1(2,4) p3(4,4) p4(7,4) p5(10,4) p7(12,4)
// 4  |       *     *        *        *     *  
// 3  |   p2(2,2)                          p8(12,2)
// 2  |       *                             *
// 1  |
// 0  -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -----> X
//      0  1  2  3  4  5  6  7  8  9  10 11 12 13 14 15

const db = [{
  id: 0,
  label: undefined,
  pos: {
    x: 2,
    y: 6
  }
}, {
  id: 1,
  label: undefined,
  pos: {
    x: 2,
    y: 4
  }
}, {
  id: 2,
  label: undefined,
  pos: {
    x: 2,
    y: 2
  }
}, {
  id: 3,
  label: undefined,
  pos: {
    x: 4,
    y: 4
  }
}, {
  id: 4,
  label: undefined,
  pos: {
    x: 7,
    y: 4
  }
}, {
  id: 5,
  label: undefined,
  pos: {
    x: 10,
    y: 4
  }
}, {
  id: 6,
  label: undefined,
  pos: {
    x: 12,
    y: 6
  }
}, {
  id: 7,
  label: undefined,
  pos: {
    x: 12,
    y: 4
  }
}, {
  id: 8,
  label: undefined,
  pos: {
    x: 12,
    y: 2
  }
}];

function dist(p1, p2) {
  let res = Math.sqrt(
    Math.pow(p2.pos.x - p1.pos.x, 2) +
    Math.pow(p2.pos.y - p1.pos.y, 2)
  );
  console.log('===>' + res);
  return res;
}

function rangeQuery(db, q, eps) {
  let neighbors = [];
  db.map(p => { // Scan all points in the database
    if (dist(p, q) <= eps) { // Compute distance and check epsilon
      neighbors.push(p); // Add to result
    }
  });
  return neighbors;
}

function DBSCAN(db, eps, minPts) {
  let c = 0; // Cluster counter
  db.map(p => {
    if (p.label) {
      return;
    } // Previously processed in inner loop
    let neighbors = rangeQuery(db, p, eps); // Find neighbors
    if (neighbors.length < minPts) { // Density check
      p.label = 'Noise'; // Label as Noise
      return;
    }
    c++; // next cluster label
    p.label = c; // Label initial point
    let seeds = neighbors.filter(n => n.id !== p.id); // Neighbors to expand
    seeds.map(s => { // Process every seed point
      if (s.label == 'Noise') {
        s.label = c;
      } // Change Noise to border point
      if (s.label) {
        return;
      } // Previously processed
      s.label = c; // Label neighbor
      let seedNeighbors = rangeQuery(db, s, eps); // Find neighbors
      if (seedNeighbors.length >= minPts) { // Density check
        seeds.push(...seedNeighbors); // Add new neighbors to seed set
      }
    });
  });
}

const eps = 2.5;
const minPts = 4;

DBSCAN(db, eps, minPts);
console.log(db);
