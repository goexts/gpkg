# GPKG

GPKG is a golang Generic Package.
This package only changes the basics of Go into a generic implementation.
This package can be used just like the base package.

## xsync.Map is a golang sync.Map with generic key and value. ##
### Benchmark xsync.Map ###
```
=== RUN   TestMapMatchesRWMutex
--- PASS: TestMapMatchesRWMutex (0.00s)
=== RUN   TestMapMatchesDeepCopy
--- PASS: TestMapMatchesDeepCopy (0.00s)
=== RUN   TestConcurrentRange
--- PASS: TestConcurrentRange (0.24s)
=== RUN   TestIssue40999
--- PASS: TestIssue40999 (0.00s)
=== RUN   TestMapRangeNestedCall
--- PASS: TestMapRangeNestedCall (0.00s)
goos: windows
goarch: amd64
pkg: github.com/ftwp/gpkg/xsync
cpu: 12th Gen Intel(R) Core(TM) i7-12700H
BenchmarkLoadMostlyHits
BenchmarkLoadMostlyHits/*xsync_test.DeepCopyMap
BenchmarkLoadMostlyHits/*xsync_test.DeepCopyMap-20         	262144190	         5.381 ns/op
BenchmarkLoadMostlyHits/*xsync_test.RWMutexMap
BenchmarkLoadMostlyHits/*xsync_test.RWMutexMap-20          	20887255	        62.18 ns/op
BenchmarkLoadMostlyHits/*sync.Map
BenchmarkLoadMostlyHits/*sync.Map-20                       	243653041	         4.795 ns/op
BenchmarkLoadMostlyMisses
BenchmarkLoadMostlyMisses/*xsync_test.DeepCopyMap
BenchmarkLoadMostlyMisses/*xsync_test.DeepCopyMap-20       	384691582	         2.974 ns/op
BenchmarkLoadMostlyMisses/*xsync_test.RWMutexMap
BenchmarkLoadMostlyMisses/*xsync_test.RWMutexMap-20        	21046132	        59.04 ns/op
BenchmarkLoadMostlyMisses/*sync.Map
BenchmarkLoadMostlyMisses/*sync.Map-20                     	319902154	         3.615 ns/op
BenchmarkLoadOrStoreBalanced
BenchmarkLoadOrStoreBalanced/*xsync_test.DeepCopyMap
    map_bench_test.go:89: DeepCopyMap has quadratic running time.
--- SKIP: BenchmarkLoadOrStoreBalanced/*xsync_test.DeepCopyMap
BenchmarkLoadOrStoreBalanced/*xsync_test.RWMutexMap
BenchmarkLoadOrStoreBalanced/*xsync_test.RWMutexMap-20     	 5228844	       223.0 ns/op
BenchmarkLoadOrStoreBalanced/*sync.Map
BenchmarkLoadOrStoreBalanced/*sync.Map-20                  	 4992765	       240.5 ns/op
BenchmarkLoadOrStoreUnique
BenchmarkLoadOrStoreUnique/*xsync_test.DeepCopyMap
    map_bench_test.go:121: DeepCopyMap has quadratic running time.
--- SKIP: BenchmarkLoadOrStoreUnique/*xsync_test.DeepCopyMap
BenchmarkLoadOrStoreUnique/*xsync_test.RWMutexMap
BenchmarkLoadOrStoreUnique/*xsync_test.RWMutexMap-20       	 3306691	       334.6 ns/op
BenchmarkLoadOrStoreUnique/*sync.Map
BenchmarkLoadOrStoreUnique/*sync.Map-20                    	 2722851	       447.8 ns/op
BenchmarkLoadOrStoreCollision
BenchmarkLoadOrStoreCollision/*xsync_test.DeepCopyMap
BenchmarkLoadOrStoreCollision/*xsync_test.DeepCopyMap-20   	694797117	         1.786 ns/op
BenchmarkLoadOrStoreCollision/*xsync_test.RWMutexMap
BenchmarkLoadOrStoreCollision/*xsync_test.RWMutexMap-20    	14281156	        70.54 ns/op
BenchmarkLoadOrStoreCollision/*sync.Map
BenchmarkLoadOrStoreCollision/*sync.Map-20                 	637767944	         2.198 ns/op
BenchmarkLoadAndDeleteBalanced
BenchmarkLoadAndDeleteBalanced/*xsync_test.DeepCopyMap
    map_bench_test.go:153: DeepCopyMap has quadratic running time.
--- SKIP: BenchmarkLoadAndDeleteBalanced/*xsync_test.DeepCopyMap
BenchmarkLoadAndDeleteBalanced/*xsync_test.RWMutexMap
BenchmarkLoadAndDeleteBalanced/*xsync_test.RWMutexMap-20   	13953342	        76.53 ns/op
BenchmarkLoadAndDeleteBalanced/*sync.Map
BenchmarkLoadAndDeleteBalanced/*sync.Map-20                	277055893	         4.906 ns/op
BenchmarkLoadAndDeleteUnique
BenchmarkLoadAndDeleteUnique/*xsync_test.DeepCopyMap
    map_bench_test.go:181: DeepCopyMap has quadratic running time.
--- SKIP: BenchmarkLoadAndDeleteUnique/*xsync_test.DeepCopyMap
BenchmarkLoadAndDeleteUnique/*xsync_test.RWMutexMap
BenchmarkLoadAndDeleteUnique/*xsync_test.RWMutexMap-20     	11954060	        86.85 ns/op
BenchmarkLoadAndDeleteUnique/*sync.Map
BenchmarkLoadAndDeleteUnique/*sync.Map-20                  	277951788	         4.794 ns/op
BenchmarkLoadAndDeleteCollision
BenchmarkLoadAndDeleteCollision/*xsync_test.DeepCopyMap
BenchmarkLoadAndDeleteCollision/*xsync_test.DeepCopyMap-20 	12151927	       101.7 ns/op
BenchmarkLoadAndDeleteCollision/*xsync_test.RWMutexMap
BenchmarkLoadAndDeleteCollision/*xsync_test.RWMutexMap-20  	17875284	        69.05 ns/op
BenchmarkLoadAndDeleteCollision/*sync.Map
BenchmarkLoadAndDeleteCollision/*sync.Map-20               	1000000000	         1.044 ns/op
BenchmarkRange
BenchmarkRange/*xsync_test.DeepCopyMap
BenchmarkRange/*xsync_test.DeepCopyMap-20                  	 1000000	      1298 ns/op
BenchmarkRange/*xsync_test.RWMutexMap
BenchmarkRange/*xsync_test.RWMutexMap-20                   	   16716	     78286 ns/op
BenchmarkRange/*sync.Map
BenchmarkRange/*sync.Map-20                                	  615364	      1774 ns/op
BenchmarkAdversarialAlloc
BenchmarkAdversarialAlloc/*xsync_test.DeepCopyMap
BenchmarkAdversarialAlloc/*xsync_test.DeepCopyMap-20       	  601627	      1798 ns/op
BenchmarkAdversarialAlloc/*xsync_test.RWMutexMap
BenchmarkAdversarialAlloc/*xsync_test.RWMutexMap-20        	11729893	       104.2 ns/op
BenchmarkAdversarialAlloc/*sync.Map
BenchmarkAdversarialAlloc/*sync.Map-20                     	 5379495	       222.2 ns/op
BenchmarkAdversarialDelete
BenchmarkAdversarialDelete/*xsync_test.DeepCopyMap
BenchmarkAdversarialDelete/*xsync_test.DeepCopyMap-20      	 9167008	       140.2 ns/op
BenchmarkAdversarialDelete/*xsync_test.RWMutexMap
BenchmarkAdversarialDelete/*xsync_test.RWMutexMap-20       	10512897	       118.8 ns/op
BenchmarkAdversarialDelete/*sync.Map
BenchmarkAdversarialDelete/*sync.Map-20                    	21112821	        54.11 ns/op
BenchmarkDeleteCollision
BenchmarkDeleteCollision/*xsync_test.DeepCopyMap
BenchmarkDeleteCollision/*xsync_test.DeepCopyMap-20        	11975760	       134.6 ns/op
BenchmarkDeleteCollision/*xsync_test.RWMutexMap
BenchmarkDeleteCollision/*xsync_test.RWMutexMap-20         	17461300	        68.87 ns/op
BenchmarkDeleteCollision/*sync.Map
BenchmarkDeleteCollision/*sync.Map-20                      	1000000000	         1.147 ns/op
PASS
```
### Benchmark sync.Map ###
```
=== RUN   TestMapMatchesRWMutex
--- PASS: TestMapMatchesRWMutex (0.00s)
=== RUN   TestMapMatchesDeepCopy
--- PASS: TestMapMatchesDeepCopy (0.00s)
=== RUN   TestConcurrentRange
--- PASS: TestConcurrentRange (0.12s)
=== RUN   TestIssue40999
--- PASS: TestIssue40999 (0.00s)
=== RUN   TestMapRangeNestedCall
--- PASS: TestMapRangeNestedCall (0.00s)
goos: windows
goarch: amd64
pkg: sync
cpu: 12th Gen Intel(R) Core(TM) i7-12700H
BenchmarkLoadMostlyHits
BenchmarkLoadMostlyHits/*sync_test.DeepCopyMap
BenchmarkLoadMostlyHits/*sync_test.DeepCopyMap-20         	237992658	         4.838 ns/op
BenchmarkLoadMostlyHits/*sync_test.RWMutexMap
BenchmarkLoadMostlyHits/*sync_test.RWMutexMap-20          	21150927	        62.42 ns/op
BenchmarkLoadMostlyHits/*sync.Map
BenchmarkLoadMostlyHits/*sync.Map-20                      	213315054	         5.118 ns/op
BenchmarkLoadMostlyMisses
BenchmarkLoadMostlyMisses/*sync_test.DeepCopyMap
BenchmarkLoadMostlyMisses/*sync_test.DeepCopyMap-20       	374251830	         3.276 ns/op
BenchmarkLoadMostlyMisses/*sync_test.RWMutexMap
BenchmarkLoadMostlyMisses/*sync_test.RWMutexMap-20        	21682214	        55.08 ns/op
BenchmarkLoadMostlyMisses/*sync.Map
BenchmarkLoadMostlyMisses/*sync.Map-20                    	358300854	         3.645 ns/op
BenchmarkLoadOrStoreBalanced
BenchmarkLoadOrStoreBalanced/*sync_test.DeepCopyMap
    map_bench_test.go:89: DeepCopyMap has quadratic running time.
--- SKIP: BenchmarkLoadOrStoreBalanced/*sync_test.DeepCopyMap
BenchmarkLoadOrStoreBalanced/*sync_test.RWMutexMap
BenchmarkLoadOrStoreBalanced/*sync_test.RWMutexMap-20     	 4767638	       243.8 ns/op
BenchmarkLoadOrStoreBalanced/*sync.Map
BenchmarkLoadOrStoreBalanced/*sync.Map-20                 	 5111434	       264.3 ns/op
BenchmarkLoadOrStoreUnique
BenchmarkLoadOrStoreUnique/*sync_test.DeepCopyMap
    map_bench_test.go:121: DeepCopyMap has quadratic running time.
--- SKIP: BenchmarkLoadOrStoreUnique/*sync_test.DeepCopyMap
BenchmarkLoadOrStoreUnique/*sync_test.RWMutexMap
BenchmarkLoadOrStoreUnique/*sync_test.RWMutexMap-20       	 2962909	       364.5 ns/op
BenchmarkLoadOrStoreUnique/*sync.Map
BenchmarkLoadOrStoreUnique/*sync.Map-20                   	 2913524	       401.4 ns/op
BenchmarkLoadOrStoreCollision
BenchmarkLoadOrStoreCollision/*sync_test.DeepCopyMap
BenchmarkLoadOrStoreCollision/*sync_test.DeepCopyMap-20   	697038572	         1.843 ns/op
BenchmarkLoadOrStoreCollision/*sync_test.RWMutexMap
BenchmarkLoadOrStoreCollision/*sync_test.RWMutexMap-20    	17120233	        65.18 ns/op
BenchmarkLoadOrStoreCollision/*sync.Map
BenchmarkLoadOrStoreCollision/*sync.Map-20                	597260366	         2.029 ns/op
BenchmarkLoadAndDeleteBalanced
BenchmarkLoadAndDeleteBalanced/*sync_test.DeepCopyMap
    map_bench_test.go:153: DeepCopyMap has quadratic running time.
--- SKIP: BenchmarkLoadAndDeleteBalanced/*sync_test.DeepCopyMap
BenchmarkLoadAndDeleteBalanced/*sync_test.RWMutexMap
BenchmarkLoadAndDeleteBalanced/*sync_test.RWMutexMap-20   	12875301	        81.73 ns/op
BenchmarkLoadAndDeleteBalanced/*sync.Map
BenchmarkLoadAndDeleteBalanced/*sync.Map-20               	252203152	         4.378 ns/op
BenchmarkLoadAndDeleteUnique
BenchmarkLoadAndDeleteUnique/*sync_test.DeepCopyMap
    map_bench_test.go:181: DeepCopyMap has quadratic running time.
--- SKIP: BenchmarkLoadAndDeleteUnique/*sync_test.DeepCopyMap
BenchmarkLoadAndDeleteUnique/*sync_test.RWMutexMap
BenchmarkLoadAndDeleteUnique/*sync_test.RWMutexMap-20     	14403375	        84.10 ns/op
BenchmarkLoadAndDeleteUnique/*sync.Map
BenchmarkLoadAndDeleteUnique/*sync.Map-20                 	280141246	         4.253 ns/op
BenchmarkLoadAndDeleteCollision
BenchmarkLoadAndDeleteCollision/*sync_test.DeepCopyMap
BenchmarkLoadAndDeleteCollision/*sync_test.DeepCopyMap-20 	11107521	       107.7 ns/op
BenchmarkLoadAndDeleteCollision/*sync_test.RWMutexMap
BenchmarkLoadAndDeleteCollision/*sync_test.RWMutexMap-20  	16917586	        66.99 ns/op
BenchmarkLoadAndDeleteCollision/*sync.Map
BenchmarkLoadAndDeleteCollision/*sync.Map-20              	879567489	         1.225 ns/op
BenchmarkRange
BenchmarkRange/*sync_test.DeepCopyMap
BenchmarkRange/*sync_test.DeepCopyMap-20                  	  767554	      1481 ns/op
BenchmarkRange/*sync_test.RWMutexMap
BenchmarkRange/*sync_test.RWMutexMap-20                   	   17650	     72680 ns/op
BenchmarkRange/*sync.Map
BenchmarkRange/*sync.Map-20                               	  951579	      1493 ns/op
BenchmarkAdversarialAlloc
BenchmarkAdversarialAlloc/*sync_test.DeepCopyMap
BenchmarkAdversarialAlloc/*sync_test.DeepCopyMap-20       	  993927	      1120 ns/op
BenchmarkAdversarialAlloc/*sync_test.RWMutexMap
BenchmarkAdversarialAlloc/*sync_test.RWMutexMap-20        	12211440	       104.6 ns/op
BenchmarkAdversarialAlloc/*sync.Map
BenchmarkAdversarialAlloc/*sync.Map-20                    	 3062185	       447.1 ns/op
BenchmarkAdversarialDelete
BenchmarkAdversarialDelete/*sync_test.DeepCopyMap
BenchmarkAdversarialDelete/*sync_test.DeepCopyMap-20      	 5875495	       183.9 ns/op
BenchmarkAdversarialDelete/*sync_test.RWMutexMap
BenchmarkAdversarialDelete/*sync_test.RWMutexMap-20       	11539803	       110.0 ns/op
BenchmarkAdversarialDelete/*sync.Map
BenchmarkAdversarialDelete/*sync.Map-20                   	22641338	        54.97 ns/op
BenchmarkDeleteCollision
BenchmarkDeleteCollision/*sync_test.DeepCopyMap
BenchmarkDeleteCollision/*sync_test.DeepCopyMap-20        	12901963	       128.5 ns/op
BenchmarkDeleteCollision/*sync_test.RWMutexMap
BenchmarkDeleteCollision/*sync_test.RWMutexMap-20         	17416570	        88.22 ns/op
BenchmarkDeleteCollision/*sync.Map
BenchmarkDeleteCollision/*sync.Map-20                     	1000000000	         1.392 ns/op
PASS
```