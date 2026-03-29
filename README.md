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

最后更新：2026-03-29 10:37:48 UTC（2026-03-29 18:37:48 UTC+8）

**代理总数：86**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 86 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 86 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 35.225.22.61:80 | ✓ 867ms | 否 | ✓ 1078ms | ✓ 1519ms | ✓ 956ms | http |
| 43.99.54.236:5555 | ✓ 656ms | ✓ 1649ms | ✓ 605ms | ✓ 795ms | ✓ 657ms | http |
| 219.117.204.211:7799 | ✓ 1526ms | ✓ 1534ms | ✓ 709ms | ✓ 1038ms | ✓ 730ms | http |
| 147.161.210.140:8800 | ✓ 1521ms | 否 | ✓ 813ms | ✓ 1616ms | ✓ 804ms | http |
| 103.84.95.54:7890 | ✓ 761ms | 否 | ✓ 978ms | ✓ 1560ms | ✓ 1777ms | http |
| 113.160.132.26:8080 | ✓ 1465ms | ✓ 1363ms | ✓ 1384ms | ✓ 1240ms | ✓ 927ms | http |
| 167.103.115.102:8800 | ✓ 1367ms | 否 | ✓ 1338ms | ✓ 1236ms | ✓ 1551ms | http |
| 167.103.34.108:8800 | ✓ 1424ms | 否 | ✓ 1380ms | ✓ 1489ms | ✓ 1333ms | http |
| 180.250.219.58:53281 | ✓ 1777ms | 否 | ✓ 1670ms | ✓ 1871ms | ✓ 1855ms | http |
| 45.140.147.82:1082 | ✓ 604ms | ✓ 1585ms | ✓ 1130ms | 否 | ✓ 1320ms | http |
| 45.140.147.82:1081 | ✓ 644ms | ✓ 1668ms | ✓ 1063ms | 否 | ✓ 1311ms | http |
| 95.213.217.168:52004 | ✓ 724ms | ✓ 1778ms | ✓ 917ms | ✓ 1793ms | 否 | http |
| 167.103.144.127:8800 | ✓ 1465ms | 否 | ✓ 1389ms | 否 | ✓ 1690ms | http |
| 192.241.132.92:80 | ✓ 664ms | ✓ 1281ms | ✓ 1343ms | ✓ 1602ms | ✓ 1555ms | http |
| 217.76.245.80:999 | ✓ 1000ms | ✓ 1471ms | ✓ 1304ms | ✓ 1885ms | ✓ 1540ms | http |
| 208.87.243.199:7878 | ✓ 376ms | ✓ 762ms | ✓ 721ms | ✓ 675ms | ✓ 574ms | http |
| 115.231.181.40:8128 | ✓ 878ms | ✓ 1619ms | ✓ 882ms | ✓ 1085ms | ✓ 892ms | http |
| 167.172.77.49:8080 | ✓ 1328ms | 否 | ✓ 1397ms | ✓ 988ms | ✓ 789ms | http |
| 167.103.31.122:8800 | ✓ 1382ms | 否 | ✓ 1303ms | ✓ 1668ms | ✓ 1499ms | http |
| 160.238.65.5:3128 | ✓ 919ms | 否 | ✓ 889ms | 否 | ✓ 1437ms | http |
| 192.73.243.65:3128 | ✓ 736ms | ✓ 1183ms | ✓ 1976ms | 否 | 否 | http |
| 47.95.231.180:8084 | 否 | ✓ 1966ms | ✓ 1191ms | ✓ 1645ms | ✓ 1644ms | http |
| 8.222.175.80:6128 | ✓ 1484ms | ✓ 1834ms | ✓ 737ms | ✓ 1015ms | ✓ 801ms | http |
| 101.43.127.100:8877 | ✓ 1742ms | ✓ 1022ms | ✓ 1165ms | ✓ 1011ms | ✓ 900ms | http |
| 147.161.239.240:8800 | ✓ 903ms | 否 | ✓ 1466ms | ✓ 1932ms | ✓ 1678ms | http |
| 120.92.212.16:7890 | ✓ 939ms | 否 | ✓ 982ms | 否 | ✓ 948ms | http |
| 177.234.217.88:999 | ✓ 1436ms | 否 | ✓ 1851ms | 否 | ✓ 1840ms | http |
| 1.231.81.166:3128 | ✓ 921ms | ✓ 923ms | ✓ 1390ms | ✓ 1232ms | ✓ 1024ms | http |
| 45.12.151.226:2829 | ✓ 944ms | ✓ 1930ms | ✓ 1399ms | 否 | ✓ 1436ms | http |
| 42.96.16.158:1311 | ✓ 1746ms | 否 | ✓ 1451ms | ✓ 1983ms | 否 | http |
| 120.132.97.88:7897 | ✓ 908ms | 否 | ✓ 871ms | ✓ 1171ms | ✓ 1484ms | http |
| 120.92.212.16:8890 | ✓ 1757ms | 否 | 否 | ✓ 1802ms | ✓ 912ms | http |
| 128.199.116.219:9090 | ✓ 1069ms | 否 | ✓ 1450ms | ✓ 1090ms | 否 | http |
| 200.125.171.254:999 | ✓ 942ms | ✓ 1651ms | ✓ 1334ms | ✓ 1480ms | ✓ 1516ms | http |
| 85.208.108.43:2094 | ✓ 353ms | 否 | ✓ 1166ms | ✓ 1244ms | ✓ 1015ms | http |
| 116.80.62.22:3128 | 否 | 否 | ✓ 1508ms | ✓ 1786ms | ✓ 1631ms | http |
| 101.32.244.83:8080 | ✓ 1382ms | 否 | ✓ 896ms | ✓ 1347ms | ✓ 1184ms | http |
| 121.43.196.210:8222 | ✓ 933ms | ✓ 1030ms | ✓ 839ms | ✓ 1075ms | ✓ 823ms | http |
| 121.43.196.213:8222 | ✓ 1045ms | ✓ 1017ms | ✓ 837ms | ✓ 1059ms | ✓ 854ms | http |
| 114.55.226.123:10086 | ✓ 1007ms | 否 | ✓ 1014ms | ✓ 1206ms | ✓ 1048ms | http |
| 59.46.216.131:30001 | ✓ 1174ms | 否 | ✓ 1007ms | 否 | ✓ 1016ms | http |
| 8.219.97.248:80 | ✓ 1686ms | 否 | ✓ 1521ms | ✓ 1396ms | 否 | http |
| 38.145.208.172:8448 | ✓ 1207ms | ✓ 1260ms | ✓ 600ms | ✓ 701ms | 否 | http |
| 129.213.162.27:17777 | ✓ 523ms | 否 | ✓ 1592ms | ✓ 1673ms | ✓ 1878ms | http |
| 103.139.138.194:3128 | 否 | 否 | ✓ 1795ms | ✓ 1633ms | ✓ 1410ms | http |
| 38.34.179.202:8449 | 否 | ✓ 849ms | ✓ 101ms | ✓ 675ms | ✓ 964ms | http |
| 38.145.208.169:8446 | 否 | 否 | ✓ 113ms | ✓ 698ms | ✓ 510ms | http |
| 38.145.208.151:8453 | 否 | ✓ 1088ms | ✓ 936ms | ✓ 914ms | ✓ 693ms | http |
| 38.34.179.67:8451 | 否 | ✓ 1727ms | ✓ 187ms | ✓ 980ms | ✓ 790ms | http |
| 38.34.179.35:8448 | 否 | 否 | ✓ 108ms | ✓ 836ms | ✓ 720ms | http |
| 45.136.130.248:8451 | 否 | ✓ 1360ms | ✓ 1263ms | ✓ 1675ms | ✓ 768ms | http |
| 38.145.208.242:8444 | 否 | ✓ 1748ms | ✓ 1033ms | 否 | ✓ 488ms | http |
| 38.145.218.87:8445 | ✓ 159ms | ✓ 621ms | ✓ 204ms | 否 | ✓ 516ms | http |
| 38.34.179.27:8451 | ✓ 189ms | ✓ 788ms | ✓ 395ms | ✓ 669ms | ✓ 617ms | http |
| 106.75.15.167:7890 | ✓ 1453ms | 否 | ✓ 1300ms | ✓ 1804ms | ✓ 1115ms | http |
| 38.34.179.39:8452 | ✓ 341ms | ✓ 1079ms | ✓ 817ms | 否 | 否 | http |
| 106.117.208.101:7890 | ✓ 970ms | ✓ 1214ms | ✓ 1033ms | ✓ 1924ms | ✓ 1964ms | http |
| 193.233.22.29:10808 | ✓ 636ms | 否 | ✓ 1563ms | ✓ 1390ms | 否 | http |
| 45.136.198.40:3128 | ✓ 860ms | ✓ 1922ms | 否 | 否 | ✓ 1980ms | http |
| 38.34.179.30:8443 | ✓ 346ms | ✓ 1705ms | ✓ 734ms | ✓ 1446ms | ✓ 1059ms | http |
| 194.67.99.223:1080 | ✓ 1296ms | ✓ 1915ms | ✓ 777ms | ✓ 1844ms | 否 | http |
| 113.255.59.226:8080 | ✓ 1133ms | 否 | ✓ 1230ms | ✓ 1045ms | ✓ 1071ms | http |
| 66.228.47.125:110 | ✓ 997ms | ✓ 1697ms | ✓ 1370ms | ✓ 1290ms | ✓ 1147ms | http |
| 103.39.51.190:8080 | ✓ 1808ms | 否 | 否 | ✓ 1508ms | ✓ 1264ms | http |
| 38.34.179.84:8451 | 否 | ✓ 1251ms | ✓ 754ms | ✓ 661ms | ✓ 662ms | http |
| 5.104.87.17:8051 | ✓ 1697ms | 否 | ✓ 1415ms | ✓ 1718ms | ✓ 1418ms | http |
| 111.198.72.198:7891 | ✓ 836ms | 否 | 否 | ✓ 1136ms | ✓ 1927ms | http |
| 38.34.179.37:8447 | ✓ 447ms | ✓ 1006ms | ✓ 803ms | ✓ 1030ms | ✓ 659ms | http |
| 103.125.16.12:8080 | ✓ 1838ms | 否 | ✓ 1556ms | ✓ 1355ms | ✓ 1313ms | http |
| 38.34.179.40:8446 | ✓ 618ms | ✓ 636ms | ✓ 256ms | ✓ 669ms | ✓ 484ms | http |
| 38.34.179.98:8453 | ✓ 1370ms | ✓ 642ms | ✓ 270ms | ✓ 960ms | ✓ 1208ms | http |
| 38.34.179.104:8446 | ✓ 1370ms | ✓ 642ms | ✓ 321ms | ✓ 1245ms | 否 | http |
| 38.34.179.101:8446 | ✓ 1370ms | ✓ 1072ms | ✓ 731ms | 否 | ✓ 699ms | http |
| 38.34.179.32:8450 | ✓ 619ms | ✓ 1057ms | ✓ 618ms | ✓ 1605ms | ✓ 477ms | http |
| 121.230.8.34:1080 | ✓ 1169ms | ✓ 1278ms | ✓ 1063ms | ✓ 1556ms | ✓ 1169ms | http |
| 121.230.8.153:1080 | ✓ 953ms | ✓ 1444ms | ✓ 1061ms | ✓ 1479ms | ✓ 1057ms | http |
| 121.230.9.248:1080 | ✓ 1650ms | 否 | ✓ 1055ms | ✓ 1384ms | ✓ 1091ms | http |
| 38.34.179.42:8445 | ✓ 618ms | 否 | ✓ 1218ms | ✓ 1260ms | ✓ 1496ms | http |
| 38.34.179.25:8444 | ✓ 1626ms | 否 | ✓ 1258ms | ✓ 1765ms | ✓ 1619ms | http |
| 38.34.179.26:8450 | ✓ 197ms | ✓ 630ms | ✓ 346ms | ✓ 764ms | ✓ 513ms | http |
| 38.145.208.239:8446 | ✓ 462ms | ✓ 1153ms | ✓ 132ms | ✓ 706ms | ✓ 527ms | http |
| 38.145.218.14:8450 | ✓ 1534ms | ✓ 771ms | ✓ 91ms | ✓ 696ms | ✓ 1364ms | http |
| 38.145.208.177:8445 | ✓ 1040ms | 否 | ✓ 1303ms | 否 | ✓ 1725ms | http |
| 38.34.179.97:8448 | 否 | 否 | ✓ 784ms | ✓ 692ms | ✓ 548ms | http |
| 38.34.179.61:8445 | ✓ 1676ms | ✓ 705ms | ✓ 323ms | ✓ 711ms | ✓ 1791ms | http |
| 121.126.185.63:25152 | ✓ 1619ms | 否 | ✓ 1337ms | 否 | ✓ 1509ms | http |

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
