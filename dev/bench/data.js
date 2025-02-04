window.BENCHMARK_DATA = {
  "lastUpdate": 1738694772210,
  "repoUrl": "https://github.com/pulsejet/YaNFD",
  "entries": {
    "Go Benchmark": [
      {
        "commit": {
          "author": {
            "email": "radialapps@gmail.com",
            "name": "Varun Patil",
            "username": "pulsejet"
          },
          "committer": {
            "email": "radialapps@gmail.com",
            "name": "Varun Patil",
            "username": "pulsejet"
          },
          "distinct": true,
          "id": "4338912afc12b7eae18097e9c579f2089eca0d2e",
          "message": "bench: limit iterations",
          "timestamp": "2025-02-04T10:19:03-08:00",
          "tree_id": "497d516392ac9b796b04bcd3042d380945eaa0d4",
          "url": "https://github.com/pulsejet/YaNFD/commit/4338912afc12b7eae18097e9c579f2089eca0d2e"
        },
        "date": 1738693234794,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkNameHash",
            "value": 370.9,
            "unit": "ns/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameStringEncode",
            "value": 8054,
            "unit": "ns/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameTlvStrEncode",
            "value": 495.6,
            "unit": "ns/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameBytesEncode",
            "value": 319.5,
            "unit": "ns/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameComponentStringEncode",
            "value": 350.6,
            "unit": "ns/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameComponentTlvStrEncode",
            "value": 29.08,
            "unit": "ns/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameClone",
            "value": 573.2,
            "unit": "ns/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameStringDecode",
            "value": 9483,
            "unit": "ns/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameTlvStrDecode",
            "value": 1039,
            "unit": "ns/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameComponentStringDecode",
            "value": 244,
            "unit": "ns/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameComponentTlvStrDecode",
            "value": 92.16,
            "unit": "ns/op",
            "extra": "150000 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "radialapps@gmail.com",
            "name": "Varun Patil",
            "username": "pulsejet"
          },
          "committer": {
            "email": "radialapps@gmail.com",
            "name": "Varun Patil",
            "username": "pulsejet"
          },
          "distinct": true,
          "id": "cd0e74b9b7d5e5ea054dfc77880184ade50a78a5",
          "message": "Add benchmem",
          "timestamp": "2025-02-04T10:30:10-08:00",
          "tree_id": "713b85c39067bf043659a672b8d40e4f727f8f4f",
          "url": "https://github.com/pulsejet/YaNFD/commit/cd0e74b9b7d5e5ea054dfc77880184ade50a78a5"
        },
        "date": 1738693851656,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkNameHash",
            "value": 367.3,
            "unit": "ns/op\t     448 B/op\t       1 allocs/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameHash - ns/op",
            "value": 367.3,
            "unit": "ns/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameHash - B/op",
            "value": 448,
            "unit": "B/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameHash - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameStringEncode",
            "value": 7953,
            "unit": "ns/op\t    3467 B/op\t      30 allocs/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameStringEncode - ns/op",
            "value": 7953,
            "unit": "ns/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameStringEncode - B/op",
            "value": 3467,
            "unit": "B/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameStringEncode - allocs/op",
            "value": 30,
            "unit": "allocs/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameTlvStrEncode",
            "value": 478.4,
            "unit": "ns/op\t     448 B/op\t       1 allocs/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameTlvStrEncode - ns/op",
            "value": 478.4,
            "unit": "ns/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameTlvStrEncode - B/op",
            "value": 448,
            "unit": "B/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameTlvStrEncode - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameBytesEncode",
            "value": 322.6,
            "unit": "ns/op\t     448 B/op\t       1 allocs/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameBytesEncode - ns/op",
            "value": 322.6,
            "unit": "ns/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameBytesEncode - B/op",
            "value": 448,
            "unit": "B/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameBytesEncode - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameComponentStringEncode",
            "value": 355.9,
            "unit": "ns/op\t      95 B/op\t       4 allocs/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameComponentStringEncode - ns/op",
            "value": 355.9,
            "unit": "ns/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameComponentStringEncode - B/op",
            "value": 95,
            "unit": "B/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameComponentStringEncode - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameComponentTlvStrEncode",
            "value": 31.05,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameComponentTlvStrEncode - ns/op",
            "value": 31.05,
            "unit": "ns/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameComponentTlvStrEncode - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameComponentTlvStrEncode - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameClone",
            "value": 557.5,
            "unit": "ns/op\t    1056 B/op\t       2 allocs/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameClone - ns/op",
            "value": 557.5,
            "unit": "ns/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameClone - B/op",
            "value": 1056,
            "unit": "B/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameClone - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameStringDecode",
            "value": 9434,
            "unit": "ns/op\t    2059 B/op\t      22 allocs/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameStringDecode - ns/op",
            "value": 9434,
            "unit": "ns/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameStringDecode - B/op",
            "value": 2059,
            "unit": "B/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameStringDecode - allocs/op",
            "value": 22,
            "unit": "allocs/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameTlvStrDecode",
            "value": 1006,
            "unit": "ns/op\t    2400 B/op\t       5 allocs/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameTlvStrDecode - ns/op",
            "value": 1006,
            "unit": "ns/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameTlvStrDecode - B/op",
            "value": 2400,
            "unit": "B/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameTlvStrDecode - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameComponentStringDecode",
            "value": 244.8,
            "unit": "ns/op\t      22 B/op\t       1 allocs/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameComponentStringDecode - ns/op",
            "value": 244.8,
            "unit": "ns/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameComponentStringDecode - B/op",
            "value": 22,
            "unit": "B/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameComponentStringDecode - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameComponentTlvStrDecode",
            "value": 91.08,
            "unit": "ns/op\t      48 B/op\t       2 allocs/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameComponentTlvStrDecode - ns/op",
            "value": 91.08,
            "unit": "ns/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameComponentTlvStrDecode - B/op",
            "value": 48,
            "unit": "B/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameComponentTlvStrDecode - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "150000 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "radialapps@gmail.com",
            "name": "Varun Patil",
            "username": "pulsejet"
          },
          "committer": {
            "email": "radialapps@gmail.com",
            "name": "Varun Patil",
            "username": "pulsejet"
          },
          "distinct": true,
          "id": "2e140707491a8943c9acc22eae39eae666f7cd5b",
          "message": "bench pr perm",
          "timestamp": "2025-02-04T10:45:31-08:00",
          "tree_id": "2cd21b9f588b17eb3f5fb39b5b4669f5e26f9877",
          "url": "https://github.com/pulsejet/YaNFD/commit/2e140707491a8943c9acc22eae39eae666f7cd5b"
        },
        "date": 1738694771887,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkNameHash",
            "value": 496.8,
            "unit": "ns/op\t     448 B/op\t       1 allocs/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameHash - ns/op",
            "value": 496.8,
            "unit": "ns/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameHash - B/op",
            "value": 448,
            "unit": "B/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameHash - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameStringEncode",
            "value": 8023,
            "unit": "ns/op\t    3467 B/op\t      30 allocs/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameStringEncode - ns/op",
            "value": 8023,
            "unit": "ns/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameStringEncode - B/op",
            "value": 3467,
            "unit": "B/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameStringEncode - allocs/op",
            "value": 30,
            "unit": "allocs/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameTlvStrEncode",
            "value": 486.9,
            "unit": "ns/op\t     448 B/op\t       1 allocs/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameTlvStrEncode - ns/op",
            "value": 486.9,
            "unit": "ns/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameTlvStrEncode - B/op",
            "value": 448,
            "unit": "B/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameTlvStrEncode - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameBytesEncode",
            "value": 307,
            "unit": "ns/op\t     448 B/op\t       1 allocs/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameBytesEncode - ns/op",
            "value": 307,
            "unit": "ns/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameBytesEncode - B/op",
            "value": 448,
            "unit": "B/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameBytesEncode - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameComponentStringEncode",
            "value": 335.1,
            "unit": "ns/op\t      95 B/op\t       4 allocs/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameComponentStringEncode - ns/op",
            "value": 335.1,
            "unit": "ns/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameComponentStringEncode - B/op",
            "value": 95,
            "unit": "B/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameComponentStringEncode - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameComponentTlvStrEncode",
            "value": 31.15,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameComponentTlvStrEncode - ns/op",
            "value": 31.15,
            "unit": "ns/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameComponentTlvStrEncode - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameComponentTlvStrEncode - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameClone",
            "value": 582.6,
            "unit": "ns/op\t    1056 B/op\t       2 allocs/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameClone - ns/op",
            "value": 582.6,
            "unit": "ns/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameClone - B/op",
            "value": 1056,
            "unit": "B/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameClone - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameStringDecode",
            "value": 9444,
            "unit": "ns/op\t    2059 B/op\t      22 allocs/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameStringDecode - ns/op",
            "value": 9444,
            "unit": "ns/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameStringDecode - B/op",
            "value": 2059,
            "unit": "B/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameStringDecode - allocs/op",
            "value": 22,
            "unit": "allocs/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameTlvStrDecode",
            "value": 1048,
            "unit": "ns/op\t    2400 B/op\t       5 allocs/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameTlvStrDecode - ns/op",
            "value": 1048,
            "unit": "ns/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameTlvStrDecode - B/op",
            "value": 2400,
            "unit": "B/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameTlvStrDecode - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameComponentStringDecode",
            "value": 245.9,
            "unit": "ns/op\t      23 B/op\t       1 allocs/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameComponentStringDecode - ns/op",
            "value": 245.9,
            "unit": "ns/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameComponentStringDecode - B/op",
            "value": 23,
            "unit": "B/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameComponentStringDecode - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameComponentTlvStrDecode",
            "value": 93.51,
            "unit": "ns/op\t      48 B/op\t       2 allocs/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameComponentTlvStrDecode - ns/op",
            "value": 93.51,
            "unit": "ns/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameComponentTlvStrDecode - B/op",
            "value": 48,
            "unit": "B/op",
            "extra": "150000 times\n4 procs"
          },
          {
            "name": "BenchmarkNameComponentTlvStrDecode - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "150000 times\n4 procs"
          }
        ]
      }
    ]
  }
}