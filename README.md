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

最后更新：2026-03-21 14:44:51 UTC（2026-03-21 22:44:51 UTC+8）

**代理总数：134**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 133 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 134 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 147.161.210.140:8800 | 否 | 否 | ✓ 896ms | ✓ 1240ms | ✓ 1126ms | http |
| 147.161.239.240:8800 | ✓ 1249ms | ✓ 1725ms | ✓ 1391ms | ✓ 1679ms | 否 | http |
| 210.76.193.248:10808 | ✓ 1124ms | ✓ 1251ms | ✓ 1250ms | 否 | ✓ 1385ms | http |
| 167.103.34.108:8800 | 否 | 否 | ✓ 1712ms | ✓ 1808ms | ✓ 1671ms | http |
| 113.160.132.26:8080 | ✓ 1912ms | ✓ 1487ms | ✓ 1264ms | ✓ 1366ms | ✓ 1068ms | http |
| 45.167.124.52:8080 | ✓ 1788ms | 否 | 否 | ✓ 1908ms | ✓ 1501ms | http |
| 38.145.208.181:8445 | ✓ 507ms | ✓ 1156ms | ✓ 419ms | ✓ 836ms | ✓ 876ms | http |
| 38.34.179.105:8449 | ✓ 516ms | 否 | ✓ 1809ms | ✓ 826ms | ✓ 888ms | http |
| 38.34.179.40:8446 | ✓ 792ms | 否 | ✓ 205ms | ✓ 1013ms | ✓ 1651ms | http |
| 38.34.179.25:8444 | ✓ 658ms | 否 | ✓ 1436ms | ✓ 895ms | ✓ 1109ms | http |
| 137.220.150.22:6005 | ✓ 860ms | 否 | 否 | ✓ 1293ms | ✓ 1162ms | http |
| 120.92.212.16:8890 | ✓ 1053ms | ✓ 1328ms | 否 | ✓ 1345ms | ✓ 1813ms | http |
| 133.242.138.34:8100 | 否 | ✓ 1564ms | 否 | ✓ 1600ms | ✓ 1516ms | http |
| 38.34.179.88:8446 | ✓ 1354ms | ✓ 770ms | 否 | ✓ 1713ms | 否 | http |
| 38.34.179.165:8446 | ✓ 329ms | ✓ 1605ms | ✓ 883ms | ✓ 1134ms | ✓ 637ms | http |
| 20.78.118.91:8561 | ✓ 1483ms | ✓ 1177ms | ✓ 547ms | ✓ 902ms | ✓ 749ms | http |
| 20.27.11.248:8561 | ✓ 1483ms | ✓ 1371ms | ✓ 545ms | ✓ 868ms | ✓ 770ms | http |
| 20.27.13.35:8561 | ✓ 1483ms | ✓ 1330ms | ✓ 548ms | ✓ 909ms | ✓ 769ms | http |
| 20.27.15.111:8561 | ✓ 1484ms | ✓ 1378ms | ✓ 542ms | ✓ 865ms | ✓ 768ms | http |
| 20.78.26.206:8561 | ✓ 1483ms | 否 | ✓ 541ms | ✓ 866ms | ✓ 699ms | http |
| 20.210.39.153:8561 | ✓ 1483ms | 否 | ✓ 555ms | ✓ 867ms | ✓ 692ms | http |
| 20.27.14.220:8561 | ✓ 1483ms | 否 | ✓ 545ms | ✓ 861ms | ✓ 720ms | http |
| 38.34.179.98:8453 | ✓ 464ms | 否 | ✓ 695ms | ✓ 1562ms | ✓ 1078ms | http |
| 35.225.22.61:80 | 否 | ✓ 1211ms | ✓ 317ms | 否 | ✓ 879ms | http |
| 142.171.224.229:7890 | ✓ 852ms | ✓ 1969ms | ✓ 299ms | ✓ 807ms | ✓ 587ms | http |
| 103.84.95.54:7890 | 否 | 否 | ✓ 913ms | ✓ 1039ms | ✓ 737ms | http |
| 167.103.31.122:8800 | ✓ 1571ms | 否 | ✓ 1375ms | ✓ 1536ms | 否 | http |
| 120.92.212.16:7890 | 否 | 否 | ✓ 1586ms | ✓ 1333ms | ✓ 1948ms | http |
| 38.34.179.172:8451 | ✓ 295ms | 否 | ✓ 231ms | ✓ 862ms | ✓ 810ms | http |
| 38.145.208.185:8449 | ✓ 415ms | 否 | ✓ 545ms | ✓ 839ms | ✓ 673ms | http |
| 139.159.99.242:8080 | ✓ 929ms | ✓ 1086ms | ✓ 922ms | ✓ 1176ms | ✓ 891ms | http |
| 219.117.204.211:7799 | ✓ 1619ms | 否 | 否 | ✓ 1294ms | ✓ 737ms | http |
| 38.34.179.54:8446 | ✓ 1252ms | ✓ 809ms | ✓ 469ms | ✓ 1237ms | ✓ 1874ms | http |
| 38.34.179.8:8449 | ✓ 826ms | 否 | ✓ 212ms | ✓ 815ms | ✓ 1982ms | http |
| 101.43.127.100:8877 | ✓ 1815ms | ✓ 1671ms | ✓ 975ms | 否 | 否 | http |
| 137.220.150.170:6005 | ✓ 953ms | 否 | ✓ 999ms | ✓ 1738ms | ✓ 932ms | http |
| 116.80.49.163:3172 | ✓ 1600ms | 否 | 否 | ✓ 1900ms | ✓ 1726ms | http |
| 137.220.150.104:6005 | 否 | 否 | ✓ 1929ms | ✓ 1282ms | ✓ 1349ms | http |
| 38.34.179.57:8453 | 否 | ✓ 959ms | ✓ 1291ms | ✓ 1611ms | ✓ 936ms | http |
| 38.34.179.173:8452 | 否 | ✓ 1023ms | ✓ 1443ms | ✓ 1071ms | ✓ 723ms | http |
| 49.156.44.114:8080 | ✓ 1865ms | 否 | ✓ 1406ms | ✓ 1531ms | 否 | http |
| 38.34.179.27:8451 | ✓ 551ms | ✓ 1480ms | ✓ 590ms | ✓ 842ms | ✓ 826ms | http |
| 148.135.116.20:8118 | ✓ 553ms | ✓ 1277ms | ✓ 733ms | 否 | 否 | http |
| 38.34.179.61:8445 | ✓ 595ms | 否 | 否 | ✓ 1725ms | ✓ 742ms | http |
| 206.81.27.105:3128 | ✓ 535ms | ✓ 1691ms | ✓ 1789ms | ✓ 1534ms | ✓ 1298ms | http |
| 137.220.150.152:6005 | ✓ 1736ms | 否 | ✓ 891ms | ✓ 1399ms | ✓ 946ms | http |
| 91.233.223.147:3128 | ✓ 897ms | ✓ 1992ms | ✓ 1168ms | 否 | ✓ 1515ms | http |
| 62.234.206.73:3128 | ✓ 1043ms | ✓ 1337ms | 否 | ✓ 1271ms | 否 | http |
| 194.67.99.223:1080 | 否 | 否 | ✓ 1967ms | ✓ 1881ms | ✓ 1401ms | http |
| 38.34.179.162:8451 | ✓ 1406ms | 否 | ✓ 735ms | ✓ 1759ms | 否 | http |
| 38.34.179.96:8451 | 否 | 否 | ✓ 310ms | ✓ 924ms | ✓ 707ms | http |
| 38.145.203.132:8452 | 否 | ✓ 968ms | ✓ 1505ms | ✓ 1984ms | ✓ 902ms | http |
| 38.34.179.86:8452 | ✓ 299ms | ✓ 896ms | ✓ 291ms | ✓ 1118ms | ✓ 727ms | http |
| 38.34.179.83:8448 | ✓ 297ms | 否 | ✓ 1656ms | ✓ 933ms | ✓ 1069ms | http |
| 45.136.130.188:8449 | ✓ 768ms | ✓ 1775ms | ✓ 1230ms | ✓ 1591ms | 否 | http |
| 45.136.130.198:8449 | ✓ 758ms | ✓ 1819ms | ✓ 1197ms | ✓ 1602ms | 否 | http |
| 45.136.130.186:8451 | ✓ 796ms | 否 | ✓ 232ms | ✓ 823ms | ✓ 836ms | http |
| 45.136.130.167:8444 | ✓ 1908ms | ✓ 1017ms | ✓ 1439ms | ✓ 1919ms | 否 | http |
| 45.136.130.169:8444 | ✓ 1908ms | ✓ 1561ms | ✓ 1778ms | ✓ 1912ms | 否 | http |
| 38.34.179.29:8443 | ✓ 255ms | ✓ 768ms | ✓ 199ms | ✓ 831ms | ✓ 645ms | http |
| 38.34.179.30:8443 | ✓ 249ms | ✓ 796ms | ✓ 202ms | ✓ 808ms | ✓ 640ms | http |
| 38.34.179.21:8443 | ✓ 262ms | ✓ 805ms | ✓ 199ms | ✓ 812ms | ✓ 631ms | http |
| 38.34.179.31:8443 | ✓ 252ms | ✓ 1333ms | ✓ 195ms | ✓ 825ms | ✓ 650ms | http |
| 91.238.105.64:2024 | ✓ 1507ms | 否 | ✓ 1570ms | 否 | ✓ 1692ms | http |
| 14.225.212.37:7890 | ✓ 1801ms | 否 | ✓ 945ms | ✓ 1201ms | ✓ 942ms | http |
| 37.187.109.70:10111 | ✓ 908ms | ✓ 1453ms | ✓ 949ms | 否 | 否 | http |
| 38.34.179.203:8451 | ✓ 864ms | ✓ 959ms | ✓ 561ms | ✓ 1964ms | ✓ 756ms | http |
| 193.23.200.251:10808 | ✓ 1365ms | ✓ 1931ms | ✓ 1493ms | 否 | 否 | http |
| 38.34.179.22:8443 | ✓ 190ms | ✓ 765ms | ✓ 193ms | ✓ 812ms | ✓ 716ms | http |
| 38.34.179.24:8443 | ✓ 205ms | ✓ 766ms | ✓ 216ms | ✓ 839ms | ✓ 642ms | http |
| 38.34.179.26:8450 | ✓ 199ms | ✓ 847ms | ✓ 202ms | ✓ 820ms | ✓ 622ms | http |
| 38.34.179.23:8444 | ✓ 192ms | ✓ 1187ms | ✓ 203ms | ✓ 795ms | ✓ 631ms | http |
| 38.34.179.28:8443 | ✓ 207ms | ✓ 1177ms | ✓ 202ms | ✓ 846ms | ✓ 613ms | http |
| 38.34.179.192:8450 | ✓ 1743ms | ✓ 880ms | ✓ 229ms | ✓ 946ms | ✓ 641ms | http |
| 38.34.179.190:8450 | ✓ 1892ms | ✓ 1469ms | ✓ 197ms | ✓ 824ms | ✓ 1266ms | http |
| 38.34.179.184:8450 | ✓ 1946ms | ✓ 1842ms | ✓ 200ms | ✓ 815ms | ✓ 1149ms | http |
| 164.90.155.209:3128 | ✓ 1493ms | ✓ 959ms | ✓ 638ms | ✓ 856ms | ✓ 627ms | http |
| 62.113.119.14:8080 | 否 | 否 | ✓ 1832ms | ✓ 1612ms | ✓ 1163ms | http |
| 114.237.77.244:1080 | ✓ 1787ms | ✓ 1349ms | ✓ 1908ms | ✓ 1624ms | ✓ 1051ms | http |
| 38.34.179.191:8453 | ✓ 209ms | ✓ 1232ms | ✓ 543ms | ✓ 1167ms | ✓ 704ms | http |
| 38.34.179.186:8444 | ✓ 203ms | ✓ 931ms | ✓ 843ms | ✓ 878ms | ✓ 794ms | http |
| 38.34.179.27:8445 | ✓ 441ms | ✓ 1865ms | ✓ 202ms | ✓ 909ms | ✓ 655ms | http |
| 38.34.179.189:8445 | ✓ 208ms | ✓ 883ms | ✓ 893ms | ✓ 849ms | ✓ 1403ms | http |
| 137.184.1.155:3128 | ✓ 372ms | ✓ 1973ms | ✓ 832ms | ✓ 833ms | ✓ 636ms | http |
| 222.184.48.241:22222 | 否 | ✓ 1336ms | ✓ 969ms | 否 | ✓ 1059ms | http |
| 183.249.5.105:22222 | 否 | 否 | ✓ 1665ms | ✓ 1721ms | ✓ 1985ms | http |
| 45.136.131.28:8444 | ✓ 262ms | 否 | ✓ 1589ms | ✓ 825ms | ✓ 945ms | http |
| 222.109.119.178:3128 | ✓ 1860ms | ✓ 1863ms | ✓ 1834ms | 否 | 否 | http |
| 166.88.55.83:7890 | ✓ 717ms | ✓ 1190ms | ✓ 706ms | ✓ 891ms | ✓ 722ms | http |
| 183.249.5.117:22222 | 否 | ✓ 1379ms | ✓ 1438ms | ✓ 1217ms | ✓ 1903ms | http |
| 202.141.161.53:30001 | ✓ 1959ms | ✓ 1496ms | ✓ 1902ms | 否 | ✓ 1191ms | http |
| 8.219.97.248:80 | ✓ 1248ms | 否 | ✓ 1594ms | ✓ 1478ms | 否 | http |
| 38.34.179.106:8446 | ✓ 847ms | ✓ 850ms | ✓ 1334ms | 否 | ✓ 706ms | http |
| 167.71.60.190:8080 | ✓ 1354ms | 否 | ✓ 1125ms | 否 | ✓ 1404ms | http |
| 59.46.216.131:30001 | ✓ 1092ms | 否 | ✓ 1191ms | ✓ 1485ms | 否 | http |
| 88.80.150.82:8080 | ✓ 1847ms | 否 | ✓ 1800ms | 否 | ✓ 1886ms | https |
| 128.199.114.189:9090 | ✓ 1102ms | 否 | ✓ 880ms | ✓ 1406ms | ✓ 957ms | http |
| 47.101.159.19:8899 | ✓ 952ms | ✓ 1146ms | ✓ 998ms | ✓ 1197ms | ✓ 939ms | http |
| 106.75.15.167:7890 | ✓ 1266ms | ✓ 1804ms | 否 | 否 | ✓ 1043ms | http |
| 5.102.109.41:999 | ✓ 721ms | ✓ 1879ms | 否 | 否 | ✓ 1008ms | http |

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
