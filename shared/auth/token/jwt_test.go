package token

import (
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const publicKey = `
-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAu1SU1LfVLPHCozMxH2Mo
4lgOEePzNm0tRgeLezV6ffAt0gunVTLw7onLRnrq0/IzW7yWR7QkrmBL7jTKEn5u
+qKhbwKfBstIs+bMY2Zkp18gnTxKLxoS2tFczGkPLPgizskuemMghRniWaoLcyeh
kd3qqGElvW/VDL5AaWTg0nLVkjRo9z+40RQzuVaE8AkAFmxZzow3x+VJYKdjykkJ
0iT9wCS0DRTXu269V264Vf/3jvredZiKRkgwlL9xNAwxXFg0x/XFw005UWVRIkdg
cKWTjpBP2dPwVZ4WWC+9aGVd+Gyn1o0CLelf4rEjGoXbAAEgAqeGUxrcIlbjXfbc
mwIDAQAB
-----END PUBLIC KEY-----
`

// token 解析 account_id
func TestJwtTokenVerifier_Verify(t *testing.T) {
	// 解析公钥
	key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKey))
	if err != nil {
		t.Fatalf("解析公钥失败: %v", err)
	}

	// 实例化 token 校验器
	v := &JwtTokenVerifier{
		PublicKey: key,
	}

	cases := []struct {
		name    string
		token   string
		now     time.Time
		want    string
		wantErr bool
	}{
		{
			name:    "合法的 token",
			token:   "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTYyNDYyMjIsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoiY29vbF9jYXIvYXV0aCIsInN1YiI6IkpvaG4gRG9lIn0.a9Z05QhqOVSDMbMirmdPtSj0kMDJpyYkru9BHI7IsAc-bEwB6siqaBbKZ9refoM7AsOUpyTctHwjBPEkxdzv6zE0TpTxDl6jeOjwDEF70fl5B0BTNqVSvINVKFwxzUlNsR1B-WrQtd5_elzkxIvMXgIyz1RJinEn-5jkdsSWBtFZMiea5Z4n1l-G-dcSBfDFnu_iQiPOgScXfiJ-VzZvf4PTCbLek2b-_ZHYrYge6kmAiV0CHJpi9Ji6ptGJohKOMz08f1w7tq6TtoA3VGP_Cs1eHulWhjCztGlqNKKf_3VJQyXghSEuoFvBIHadSet2pfjZsyAM9YqJnqjcr4cj1w",
			now:     time.Unix(1516239122, 0),
			want:    "John Doe",
			wantErr: false,
		},
		{
			name: "非法(伪造)的 token [ 错误的签名 ]",
			//token:   "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTYyNDYyMjIsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoiY29vbF9jYXIvYXV0aCIsInN1YiI6IkpvaG4gRG9lMSJ9.sJDNEZfCPHobtshBs0wcKCwqcfXv-gK92dt7WfN0eVveM2VAZ5z3twh5BlwiAu4hvifom0wpacNXE_lS1ekk1ahhGEaQpqPpE03-KTvOh3ngH4Fou5hQgplhJNwACkKyGMRKmitihxCOQ7DWvhebf1W6tP4Mv7on5WYa4AsFa-ZOmCVgIfdizClBgMh4DiUiVdXAJq-rtg2nlTlUV-0qiXlZKdYHRtjDSKkbRDcMW9rKCDqN0bd5ZM-4D6uxOxDihEjWwWxBTrbuIFFzWx-pTabN4PYS4rNodtnbD3iv2rYrxI6rrbeOQw4DSM5RxcY7rLVrD0elVguhbv_gtDOmCA",
			token:   "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTYyNDYyMjIsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoiY29vbF9jYXIvYXV0aCIsInN1YiI6IkpvaG4gRG9lMSJ9.a9Z05QhqOVSDMbMirmdPtSj0kMDJpyYkru9BHI7IsAc-bEwB6siqaBbKZ9refoM7AsOUpyTctHwjBPEkxdzv6zE0TpTxDl6jeOjwDEF70fl5B0BTNqVSvINVKFwxzUlNsR1B-WrQtd5_elzkxIvMXgIyz1RJinEn-5jkdsSWBtFZMiea5Z4n1l-G-dcSBfDFnu_iQiPOgScXfiJ-VzZvf4PTCbLek2b-_ZHYrYge6kmAiV0CHJpi9Ji6ptGJohKOMz08f1w7tq6TtoA3VGP_Cs1eHulWhjCztGlqNKKf_3VJQyXghSEuoFvBIHadSet2pfjZsyAM9YqJnqjcr4cj1w",
			now:     time.Unix(1516239122, 0),
			want:    "",
			wantErr: true,
		},
		{
			name:    "过期的 token",
			token:   "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTYyNDYyMjIsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoiY29vbF9jYXIvYXV0aCIsInN1YiI6IkpvaG4gRG9lIn0.a9Z05QhqOVSDMbMirmdPtSj0kMDJpyYkru9BHI7IsAc-bEwB6siqaBbKZ9refoM7AsOUpyTctHwjBPEkxdzv6zE0TpTxDl6jeOjwDEF70fl5B0BTNqVSvINVKFwxzUlNsR1B-WrQtd5_elzkxIvMXgIyz1RJinEn-5jkdsSWBtFZMiea5Z4n1l-G-dcSBfDFnu_iQiPOgScXfiJ-VzZvf4PTCbLek2b-_ZHYrYge6kmAiV0CHJpi9Ji6ptGJohKOMz08f1w7tq6TtoA3VGP_Cs1eHulWhjCztGlqNKKf_3VJQyXghSEuoFvBIHadSet2pfjZsyAM9YqJnqjcr4cj1w",
			now:     time.Unix(1616239122, 0),
			want:    "",
			wantErr: true,
		},
		{
			name:    "错误(格式)的 token",
			token:   "bad token",
			now:     time.Unix(1516239122, 0),
			want:    "",
			wantErr: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			jwt.TimeFunc = func() time.Time {
				return c.now
			}
			accountID, err := v.Verify(c.token)

			// 正确案例 && 有错
			if !c.wantErr && err != nil {
				t.Errorf("正确案例校验失败: %v", err)
			}

			// 错误案例 && 没哟错误信息
			if c.wantErr && err == nil {
				t.Errorf("错误案例没有错误信息")
			}

			// 结果不对
			if accountID != c.want {
				t.Errorf("错误的 account_id, 期望值: %q; 实际值: %q", c.want, accountID)
			}
		})
	}

	token := "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTYyNDYyMjIsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoiY29vbF9jYXIvYXV0aCIsInN1YiI6IkpvaG4gRG9lIn0.a9Z05QhqOVSDMbMirmdPtSj0kMDJpyYkru9BHI7IsAc-bEwB6siqaBbKZ9refoM7AsOUpyTctHwjBPEkxdzv6zE0TpTxDl6jeOjwDEF70fl5B0BTNqVSvINVKFwxzUlNsR1B-WrQtd5_elzkxIvMXgIyz1RJinEn-5jkdsSWBtFZMiea5Z4n1l-G-dcSBfDFnu_iQiPOgScXfiJ-VzZvf4PTCbLek2b-_ZHYrYge6kmAiV0CHJpi9Ji6ptGJohKOMz08f1w7tq6TtoA3VGP_Cs1eHulWhjCztGlqNKKf_3VJQyXghSEuoFvBIHadSet2pfjZsyAM9YqJnqjcr4cj1w"
	jwt.TimeFunc = func() time.Time {
		return time.Unix(1516239122, 0)
	}
	accountID, err := v.Verify(token)
	if err != nil {
		t.Fatalf("获取 account_id 失败: %v", err)
	}

	want := "John Doe"
	if accountID != want {
		t.Fatalf("错误的 account_id, 期望值: %q; 实际值: %q ", want, accountID)
	}
}
