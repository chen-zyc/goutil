package hash

func RSHash(key string) uint64 {
	// 从Robert Sedgwicks的 Algorithms in C一书中得到了
	a, b := uint64(63689), uint64(378551)
	hash := uint64(0)
	for i, n := 0, len(key); i < n; i++ {
		hash = hash*a + uint64(key[i])
		a = a * b
	}
	return hash
}

func JSHash(key string) uint64 {
	// Justin Sobel写的一个位操作的哈希函数
	hash := uint64(1315423911)
	for i, n := 0, len(key); i < n; i++ {
		hash ^= hash<<5 + uint64(key[i]) + hash>>2
	}
	return hash
}

func PJWHash(key string) uint64 {
	// 该散列算法是基于贝尔实验室的彼得J温伯格的的研究。在Compilers一书中（原则，技术和工具），建议采用这个算法的散列函数的哈希方法
	const (
		bitsInUint    uint64 = 4 * 8
		threeQuarters uint64 = bitsInUint * 3 / 4
		oneEighth     uint64 = bitsInUint / 8
		highBits      uint64 = 0xFFFFFFFF << (bitsInUint - oneEighth)
	)
	var (
		hash uint64 = 0
		test uint64 = 0
	)
	for i, n := 0, len(key); i < n; i++ {
		hash = hash<<oneEighth + uint64(key[i])
		if test = hash & highBits; test != 0 {
			hash = (hash ^ (test >> threeQuarters)) & (^highBits)
		}
	}
	return hash
}

func ELFHash(key string) uint64 {
	// 和PJW很相似，在Unix系统中使用的较多
	hash, x := uint64(0), uint64(0)
	for i, n := 0, len(key); i < n; i++ {
		hash = hash<<4 + uint64(key[i])
		if x = hash & 0xF0000000; x != 0 {
			hash ^= x >> 24
		}
		hash &= ^x
	}
	return hash
}

func BKDRHash(key string) uint64 {
	// 这个算法来自Brian Kernighan 和 Dennis Ritchie的 The C Programming Language。
	// 这是一个很简单的哈希算法,使用了一系列奇怪的数字,形式如31,3131,31...31,看上去和DJB算法很相似。
	// 这个就是Java的字符串哈希函数
	const seed = 131 // 31 131 1313 13131 131313 etc..
	hash := uint64(0)
	for i, n := 0, len(key); i < n; i++ {
		hash = hash*seed + uint64(key[i])
	}
	return hash
}

func SDBMHash(key string) uint64 {
	// 这个算法在开源的SDBM中使用，似乎对很多不同类型的数据都能得到不错的分布。
	hash := uint64(0)
	for i, n := 0, len(key); i < n; i++ {
		hash = uint64(key[i]) + hash<<6 + hash<<16 - hash
	}
	return hash
}

func DJBHash(key string) uint64 {
	// 这个算法是Daniel J.Bernstein 教授发明的，是目前公布的最有效的哈希函数。
	hash := uint64(5381)
	for i, n := 0, len(key); i < n; i++ {
		hash = hash<<5 + hash + uint64(key[i])
	}
	return hash
}

func DEKHash(key string) uint64 {
	// 由伟大的Knuth在《编程的艺术 第三卷》的第六章排序和搜索中给出。
	hash := uint64(len(key))
	for i, n := 0, len(key); i < n; i++ {
		hash = (hash << 5) ^ (hash >> 27) ^ uint64(key[i])
	}
	return hash
}

func APHash(key string) uint64 {
	// Arash Partow发明的一种hash算法, 比较优秀的一种哈希算法
	hash := uint64(0)
	for i, n := 0, len(key); i < n; i++ {
		if i&1 == 0 {
			hash ^= (hash << 7) ^ uint64(key[i]) ^ (hash >> 3)
		} else {
			hash ^= ^((hash << 11) ^ uint64(key[i]) ^ (hash >> 5))
		}
	}
	return hash
}

func FNVHash(key string) uint64 {
	// Unix system系统中使用的一种著名hash算法，后来微软也在其hash_map中实现。
	if key == "" {
		return 0
	}
	hash := uint64(2166136261)
	for i, n := 0, len(key); i < n; i++ {
		hash *= 16777619
		hash ^= uint64(key[i])
	}
	return hash
}

func MYSQLHash(key string) uint64 {
	// MySQL中出现的字符串哈希函数
	nr, nr2 := uint64(1), uint64(4)
	for i, n := 0, len(key); i < n; i++ {
		nr ^= ((nr&63)+nr2)*uint64(key[i]) + (nr << 8)
		nr2 += 3
	}
	return nr
}
