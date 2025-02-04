window.BENCHMARK_DATA = {
  "lastUpdate": 1738693235192,
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
      }
    ]
  }
}