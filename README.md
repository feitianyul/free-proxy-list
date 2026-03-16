**中文** | [English (README_EN.md)](README_EN.md)

---

<p align="center">
  <img src="https://img.shields.io/badge/每1小时更新-通过-success">  
  <br>
  <img src="https://img.shields.io/website/https/getfreeproxy.com.svg">
  <img src="https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/total.svg">
  <img src="https://img.shields.io/github/last-commit/feitianyul/free-proxy-list.svg">
  <img src="https://img.shields.io/github/license/feitianyul/free-proxy-list.svg">
  
  <br>
  <br>
  <a href="https://getfreeproxy.com/lists/" title="可用代理列表">可用代理列表</a> | <a href="https://getfreeproxy.com/tools/proxy-checker" title="在线代理检测">免费代理检测</a> | <a href="https://getfreeproxy.com/tools/proxy-protocol-parser" title="代理协议解析">通用代理协议解析</a> | <a href="https://developer.getfreeproxy.com/" title="代理 API">免费代理 API</a>
  <br>
</p>

# 🌎 GetFreeProxy (GFP)：免费代理列表

**GetFreeProxy (GFP)** 是一个开源项目，自动从互联网聚合并校验免费代理，旨在为开发者、研究人员及需要代理服务的用户提供新鲜、可靠、可用的公共代理列表。

列表按小时更新，确保您始终能获取到最新的可用代理。

---

## 📖 项目说明

本项目为开源免费代理聚合与校验工具，从互联网公开源拉取代理并**仅保留 HTTP、HTTPS** 两种类型，经校验后生成列表，供开发者、研究人员等使用。

### 本仓库特点

- **仅保留两种代理**：HTTP、HTTPS，不收录 SOCKS、VMess、Trojan、VLESS、SS/SSR、Hysteria 等其它协议。
- **校验规则**：五域名中**任意 3 个**在 2 秒内成功（HTTP 200）即视为该协议通过。对每条代理访问以下五个地址验证（优先 HEAD，不支持则回退 GET）：
  - `https://www.eastmoney.com/`
  - `https://www.sse.com.cn/`
  - `https://finance.sina.com.cn/`（新浪财经）
  - `https://web.ifzq.gtimg.cn/`
  - `https://proxy.finance.qq.com/`
  每个代理分别以 **HTTP 代理** 和 **HTTPS 代理** 各测一次；**协议** 写入 meta：只通 HTTP→`http`，只通 HTTPS→`https`，两个都通→`http/s`。去重按 **协议+IP+端口**。校验时**多代理并发**、**单代理内五域名并行**。列表直显：表格列为「代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议」。
- **更新频率**：列表按小时更新，保证可用代理的时效性。
- **并发参数**：校验 worker 数可通过 `-check-workers`（如 `-check-workers=4000`）或环境变量 `GFP_CHECK_WORKERS` 设置，默认 4000，最大 4000。遇目标站限流可适当调低。

### 工作流程

1. **拉取**：从 `sources/` 目录下配置的源（仅处理 `http.txt`、`https.txt`）拉取原始代理数据，支持动态 URL 及 Base64 等格式。
2. **解析与规范化**：将原始数据解析为标准代理格式（协议、IP、端口、认证等）。
3. **校验**：对 HTTP/HTTPS 代理通过上述验证与 2 秒超时规则进行筛选。
4. **去重与存储**：通过校验的代理去重后写入内存。
5. **生成列表**：按协议生成 `list/` 目录下的 `http.txt`、`https.txt`，并更新统计与 README 中的下载表格。

自动化由 GitHub Actions 执行：**全量流程**（抓取→解析→验证→生成列表）**每 6 小时**运行一次；**轻量复测**（对已有列表做连通性复测、剔除失效代理）**每 1 小时**运行一次。全量任务最长运行 12 小时，超时才会取消。下表「最后更新」时间为 UTC 及 UTC+8。

### 支持的代理格式示例

| 类型 | 格式 | 示例 |
| :--- | :--- | :--- |
| **HTTP/HTTPS** | `http://ip:port` | `http://1.2.3.4:8080` |
| | `https://ip:port` | `https://1.2.3.4:8080` |
| | `http://user:pass@ip:port` | `http://user:pass@1.2.3.4:8080` |

---

## 🔗 直接下载链接

点击下方表格中您需要的协议类型即可获取最新列表，链接始终指向最近更新的代理文件。上述三个文件（http.txt、https.txt、passed.txt）同时发布至 [GitHub Releases (tag: lists)](https://github.com/feitianyul/free-proxy-list/releases/tag/lists)，每次运行覆盖同版本附件，可固定使用该 Release 的附件 URL。

<!-- BEGIN PROXY LIST -->

最后更新：2026-03-16 08:15:08 UTC（2026-03-16 16:15:08 UTC+8）

**代理总数：176**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 175 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 176 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 8.209.239.31:30000 | ✓ 1178ms | 否 | ✓ 703ms | ✓ 967ms | ✓ 1306ms | http |
| 202.155.12.161:443 | ✓ 1178ms | 否 | ✓ 1217ms | ✓ 1339ms | ✓ 1219ms | http |
| 47.79.40.38:55000 | ✓ 1183ms | 否 | ✓ 1460ms | ✓ 977ms | ✓ 839ms | http |
| 219.117.204.211:7799 | ✓ 1182ms | 否 | ✓ 1374ms | ✓ 1982ms | 否 | http |
| 109.73.195.10:8888 | ✓ 1502ms | 否 | ✓ 1721ms | 否 | ✓ 1987ms | http |
| 113.160.132.26:8080 | ✓ 1936ms | 否 | ✓ 1576ms | ✓ 1470ms | ✓ 1172ms | http |
| 1.231.81.166:3128 | ✓ 928ms | ✓ 1167ms | ✓ 1574ms | ✓ 1296ms | ✓ 1141ms | http |
| 38.180.2.107:3128 | ✓ 972ms | ✓ 1967ms | ✓ 1899ms | 否 | ✓ 1975ms | http |
| 211.171.114.154:3128 | 否 | ✓ 1614ms | ✓ 1674ms | ✓ 1973ms | ✓ 1252ms | http |
| 217.76.245.80:999 | ✓ 807ms | 否 | ✓ 1054ms | ✓ 1446ms | ✓ 1158ms | http |
| 38.34.179.14:8450 | 否 | ✓ 1981ms | 否 | ✓ 959ms | ✓ 730ms | http |
| 35.225.22.61:80 | ✓ 750ms | 否 | ✓ 1587ms | 否 | ✓ 1147ms | http |
| 104.129.202.127:12354 | ✓ 402ms | ✓ 1124ms | ✓ 322ms | ✓ 1045ms | ✓ 786ms | http |
| 104.129.202.127:10810 | ✓ 402ms | ✓ 1116ms | ✓ 349ms | ✓ 1065ms | ✓ 919ms | http |
| 120.232.242.124:22222 | 否 | 否 | ✓ 1162ms | ✓ 1450ms | ✓ 1126ms | http |
| 120.92.212.16:7890 | 否 | 否 | ✓ 1166ms | ✓ 1730ms | ✓ 1211ms | http |
| 120.92.212.16:8890 | 否 | 否 | ✓ 1160ms | ✓ 1508ms | ✓ 1145ms | http |
| 101.43.127.100:8877 | 否 | ✓ 1911ms | ✓ 1103ms | ✓ 1371ms | ✓ 1050ms | http |
| 81.70.169.194:80 | ✓ 1122ms | 否 | ✓ 1267ms | ✓ 1532ms | 否 | http |
| 38.145.203.19:8447 | ✓ 1007ms | 否 | ✓ 877ms | ✓ 1336ms | ✓ 1102ms | http |
| 43.153.28.68:3128 | ✓ 944ms | 否 | ✓ 484ms | ✓ 1002ms | 否 | http |
| 38.145.220.33:8448 | ✓ 1021ms | 否 | ✓ 1586ms | ✓ 917ms | ✓ 932ms | http |
| 38.34.179.60:8450 | ✓ 1761ms | ✓ 1574ms | ✓ 1122ms | ✓ 1906ms | ✓ 770ms | http |
| 38.34.179.203:8450 | ✓ 1256ms | 否 | ✓ 1192ms | ✓ 1974ms | 否 | http |
| 38.145.203.39:8450 | ✓ 312ms | ✓ 1008ms | ✓ 340ms | ✓ 990ms | ✓ 720ms | http |
| 38.145.218.229:8450 | ✓ 798ms | ✓ 915ms | ✓ 376ms | ✓ 971ms | ✓ 879ms | http |
| 38.145.220.198:8450 | ✓ 607ms | ✓ 903ms | ✓ 615ms | ✓ 1122ms | ✓ 840ms | http |
| 38.145.220.36:8450 | ✓ 922ms | ✓ 947ms | ✓ 876ms | ✓ 966ms | ✓ 755ms | http |
| 38.145.203.97:8448 | ✓ 334ms | ✓ 938ms | ✓ 1476ms | ✓ 1084ms | ✓ 719ms | http |
| 38.145.203.106:8448 | ✓ 352ms | ✓ 928ms | ✓ 1467ms | ✓ 1085ms | ✓ 740ms | http |
| 45.136.130.196:8450 | ✓ 1150ms | ✓ 929ms | ✓ 665ms | ✓ 1074ms | ✓ 905ms | http |
| 38.34.178.111:8450 | ✓ 743ms | 否 | ✓ 988ms | ✓ 924ms | ✓ 762ms | http |
| 38.34.179.35:8448 | ✓ 384ms | 否 | ✓ 690ms | ✓ 1932ms | ✓ 803ms | http |
| 45.136.130.216:8450 | ✓ 732ms | ✓ 1822ms | ✓ 1109ms | ✓ 957ms | ✓ 1187ms | http |
| 117.159.239.87:22222 | 否 | ✓ 1366ms | ✓ 968ms | ✓ 1297ms | ✓ 1110ms | http |
| 147.161.210.140:8800 | ✓ 854ms | ✓ 1420ms | ✓ 844ms | ✓ 1247ms | ✓ 1618ms | http |
| 45.136.130.160:8447 | ✓ 1529ms | ✓ 904ms | ✓ 795ms | ✓ 1930ms | ✓ 888ms | http |
| 117.159.239.46:22222 | ✓ 1023ms | ✓ 1300ms | ✓ 1030ms | ✓ 1327ms | ✓ 1043ms | http |
| 45.93.31.93:6005 | ✓ 1907ms | 否 | 否 | ✓ 1913ms | ✓ 1265ms | http |
| 162.240.154.26:3128 | ✓ 829ms | 否 | ✓ 1543ms | ✓ 999ms | 否 | http |
| 45.136.130.223:8443 | ✓ 414ms | ✓ 1828ms | ✓ 338ms | ✓ 996ms | ✓ 743ms | http |
| 45.136.130.210:8448 | ✓ 433ms | ✓ 1991ms | ✓ 332ms | ✓ 927ms | ✓ 751ms | http |
| 38.145.220.145:8447 | ✓ 857ms | 否 | ✓ 358ms | ✓ 962ms | ✓ 743ms | http |
| 38.145.220.137:8447 | ✓ 878ms | 否 | ✓ 341ms | ✓ 984ms | ✓ 742ms | http |
| 45.136.130.212:8448 | ✓ 429ms | ✓ 1308ms | ✓ 580ms | ✓ 1986ms | ✓ 728ms | http |
| 45.136.130.213:8448 | ✓ 461ms | ✓ 944ms | ✓ 902ms | ✓ 1399ms | ✓ 1770ms | http |
| 168.235.110.63:3128 | ✓ 222ms | ✓ 1912ms | ✓ 765ms | ✓ 1161ms | 否 | http |
| 117.159.239.49:22222 | ✓ 942ms | ✓ 1307ms | ✓ 973ms | 否 | ✓ 1070ms | http |
| 120.238.159.250:22222 | 否 | ✓ 1516ms | ✓ 1146ms | 否 | ✓ 1084ms | http |
| 133.242.138.34:8100 | ✓ 1391ms | ✓ 1222ms | ✓ 808ms | ✓ 1241ms | ✓ 980ms | http |
| 120.238.159.189:22222 | ✓ 1141ms | ✓ 1452ms | 否 | ✓ 1352ms | ✓ 1063ms | http |
| 85.198.96.242:3128 | ✓ 981ms | 否 | 否 | ✓ 1605ms | ✓ 1247ms | http |
| 101.43.255.96:80 | ✓ 1188ms | ✓ 1580ms | ✓ 1180ms | ✓ 1593ms | ✓ 1174ms | http |
| 103.84.95.54:7890 | ✓ 1193ms | 否 | ✓ 1063ms | ✓ 1940ms | 否 | http |
| 137.220.150.104:6005 | ✓ 1510ms | 否 | ✓ 1111ms | ✓ 1726ms | ✓ 1543ms | http |
| 137.220.151.110:6005 | ✓ 1508ms | 否 | ✓ 1829ms | ✓ 1970ms | ✓ 1314ms | http |
| 38.145.208.172:8448 | ✓ 354ms | ✓ 910ms | ✓ 952ms | ✓ 1154ms | ✓ 717ms | http |
| 38.145.208.210:8448 | ✓ 495ms | ✓ 1436ms | ✓ 288ms | ✓ 1120ms | ✓ 722ms | http |
| 38.145.218.189:8447 | ✓ 365ms | ✓ 1773ms | ✓ 306ms | ✓ 1034ms | ✓ 720ms | http |
| 38.145.220.149:8448 | ✓ 362ms | ✓ 1601ms | ✓ 295ms | ✓ 1217ms | ✓ 700ms | http |
| 45.136.130.238:8447 | ✓ 292ms | ✓ 890ms | ✓ 1651ms | ✓ 945ms | ✓ 707ms | http |
| 45.136.130.241:8448 | ✓ 292ms | 否 | ✓ 946ms | ✓ 944ms | ✓ 696ms | http |
| 38.145.208.135:8448 | ✓ 301ms | ✓ 909ms | ✓ 295ms | ✓ 1050ms | ✓ 691ms | http |
| 45.136.130.244:8448 | ✓ 919ms | ✓ 947ms | ✓ 289ms | ✓ 930ms | ✓ 696ms | http |
| 38.145.208.136:8448 | ✓ 316ms | ✓ 942ms | ✓ 312ms | ✓ 971ms | ✓ 702ms | http |
| 45.136.130.245:8447 | ✓ 897ms | ✓ 956ms | ✓ 305ms | ✓ 928ms | ✓ 935ms | http |
| 38.145.208.160:8447 | ✓ 881ms | ✓ 979ms | ✓ 820ms | ✓ 1480ms | ✓ 1121ms | http |
| 45.136.130.240:8448 | ✓ 415ms | ✓ 1682ms | ✓ 359ms | ✓ 893ms | ✓ 731ms | http |
| 38.145.208.201:8447 | ✓ 983ms | ✓ 1438ms | ✓ 1793ms | ✓ 1403ms | ✓ 898ms | http |
| 120.198.141.79:22222 | ✓ 1139ms | ✓ 1484ms | 否 | 否 | ✓ 1086ms | http |
| 38.145.208.144:8443 | ✓ 596ms | 否 | ✓ 335ms | ✓ 983ms | ✓ 724ms | http |
| 38.34.179.78:8448 | ✓ 406ms | ✓ 1666ms | ✓ 301ms | ✓ 934ms | ✓ 702ms | http |
| 38.145.218.212:8448 | ✓ 1175ms | ✓ 1928ms | ✓ 426ms | ✓ 897ms | ✓ 878ms | http |
| 45.136.130.243:8448 | ✓ 352ms | ✓ 1694ms | ✓ 291ms | ✓ 1692ms | ✓ 1702ms | http |
| 213.219.214.45:443 | ✓ 722ms | 否 | ✓ 788ms | 否 | ✓ 1314ms | http |
| 38.145.208.95:8443 | ✓ 1318ms | ✓ 988ms | ✓ 1106ms | ✓ 1429ms | ✓ 902ms | http |
| 38.145.208.142:8444 | ✓ 618ms | 否 | ✓ 318ms | ✓ 938ms | ✓ 1156ms | http |
| 45.88.0.99:3128 | ✓ 860ms | ✓ 1551ms | 否 | 否 | ✓ 927ms | http |
| 45.88.0.116:3128 | ✓ 861ms | 否 | ✓ 1889ms | ✓ 1739ms | ✓ 919ms | http |
| 45.88.0.113:3128 | ✓ 860ms | 否 | ✓ 1888ms | ✓ 1742ms | ✓ 916ms | http |
| 45.88.0.111:3128 | ✓ 863ms | 否 | ✓ 1887ms | ✓ 1737ms | ✓ 917ms | http |
| 45.88.0.114:3128 | ✓ 860ms | 否 | ✓ 1887ms | ✓ 1742ms | ✓ 917ms | http |
| 45.88.0.98:3128 | ✓ 861ms | ✓ 1660ms | 否 | ✓ 1967ms | ✓ 914ms | http |
| 38.145.218.232:8448 | ✓ 1254ms | ✓ 1797ms | ✓ 415ms | ✓ 957ms | ✓ 1445ms | http |
| 38.145.208.188:8447 | ✓ 360ms | ✓ 1432ms | ✓ 1415ms | ✓ 1230ms | ✓ 746ms | http |
| 38.145.220.49:8447 | ✓ 396ms | 否 | ✓ 1484ms | ✓ 908ms | ✓ 1030ms | http |
| 45.136.131.29:8450 | ✓ 393ms | ✓ 1815ms | ✓ 506ms | ✓ 1498ms | ✓ 683ms | http |
| 38.34.179.124:8450 | ✓ 1085ms | ✓ 1780ms | ✓ 344ms | ✓ 1101ms | ✓ 814ms | http |
| 38.34.183.224:8448 | ✓ 1820ms | ✓ 1125ms | ✓ 347ms | ✓ 1127ms | ✓ 1683ms | http |
| 160.250.4.245:1 | ✓ 1171ms | 否 | ✓ 1914ms | 否 | ✓ 1333ms | http |
| 45.136.130.211:8447 | ✓ 1859ms | ✓ 1329ms | ✓ 474ms | ✓ 1101ms | ✓ 822ms | http |
| 45.136.131.54:8448 | ✓ 1206ms | 否 | ✓ 463ms | ✓ 1400ms | ✓ 1766ms | http |
| 222.184.48.252:22222 | 否 | ✓ 1380ms | ✓ 1166ms | ✓ 1451ms | ✓ 1110ms | http |
| 43.143.108.2:9701 | ✓ 1082ms | 否 | 否 | ✓ 1371ms | ✓ 1095ms | http |
| 38.145.218.101:8447 | ✓ 1850ms | ✓ 1474ms | ✓ 1850ms | ✓ 1555ms | ✓ 1369ms | http |
| 222.184.48.235:22222 | 否 | ✓ 1270ms | ✓ 1417ms | 否 | ✓ 1533ms | http |
| 45.136.130.198:8443 | ✓ 1039ms | ✓ 1422ms | ✓ 1996ms | ✓ 1421ms | 否 | http |
| 38.145.218.228:8447 | ✓ 829ms | 否 | ✓ 323ms | ✓ 926ms | 否 | http |
| 38.145.208.179:8447 | ✓ 810ms | 否 | ✓ 354ms | ✓ 967ms | ✓ 1337ms | http |
| 45.136.130.177:8448 | ✓ 1304ms | 否 | ✓ 683ms | 否 | ✓ 756ms | http |

<!-- END PROXY TABLE -->

## 🤝 参与贡献

本项目由社区驱动，欢迎任何形式的贡献。最简单的参与方式就是添加新的代理数据源。

请先阅读 **[贡献指南](CONTRIBUTING.md)** 了解如何开始。

## 🙏 支持本项目

如果您觉得本项目有帮助，欢迎给予支持，让更多人看到并参与贡献。

-   在 GitHub 上 **给本仓库加星** ⭐️
-   **分享**给朋友和同事

## ⚠️ 免责声明

-   本仓库中的代理均来自公开来源，不保证其速度、安全性或可用性。
-   使用这些代理的风险由您自行承担。
-   本仓库维护者不对任何滥用行为负责。请勿将代理用于非法用途。

## 📝 许可证

本仓库采用 MIT 许可证发布。详见 [LICENSE](LICENSE)。

## Stars
[![Star History Chart](https://api.star-history.com/svg?repos=feitianyul/free-proxy-list&type=Date)](https://star-history.com/#feitianyul/free-proxy-list&Date)
