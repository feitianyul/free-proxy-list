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

最后更新：2026-03-13 06:40:00 UTC（2026-03-13 14:40:00 UTC+8）

**代理总数：85**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 85 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 85 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 320ms | ✓ 1748ms | ✓ 1042ms | ✓ 1295ms | 否 | http |
| 1.231.81.166:3128 | ✓ 1731ms | ✓ 1671ms | ✓ 992ms | ✓ 1235ms | ✓ 788ms | http |
| 113.160.132.26:8080 | ✓ 1810ms | ✓ 1471ms | ✓ 1464ms | ✓ 1406ms | ✓ 1265ms | http |
| 193.168.173.136:443 | ✓ 1451ms | 否 | ✓ 1331ms | 否 | ✓ 1542ms | http |
| 45.167.124.52:8080 | 否 | ✓ 1819ms | ✓ 676ms | 否 | ✓ 1797ms | http |
| 120.92.212.16:7890 | ✓ 977ms | ✓ 1163ms | ✓ 919ms | ✓ 1178ms | ✓ 922ms | http |
| 120.92.212.16:8890 | 否 | ✓ 1406ms | ✓ 1135ms | ✓ 1402ms | ✓ 1154ms | http |
| 152.42.213.210:8080 | ✓ 1733ms | 否 | ✓ 1391ms | ✓ 1220ms | ✓ 1302ms | http |
| 115.231.181.40:8128 | ✓ 908ms | ✓ 928ms | ✓ 943ms | 否 | ✓ 815ms | http |
| 150.107.140.238:3128 | 否 | 否 | ✓ 1662ms | ✓ 1339ms | ✓ 1271ms | http |
| 113.59.32.160:22222 | ✓ 1220ms | ✓ 1427ms | ✓ 1038ms | 否 | ✓ 932ms | http |
| 120.232.242.119:22222 | ✓ 1020ms | ✓ 1330ms | ✓ 960ms | ✓ 1252ms | 否 | http |
| 117.159.239.49:22222 | 否 | ✓ 1162ms | ✓ 863ms | ✓ 1218ms | ✓ 944ms | http |
| 45.136.131.63:8443 | 否 | ✓ 1904ms | 否 | ✓ 952ms | ✓ 646ms | http |
| 81.70.169.194:80 | ✓ 1188ms | 否 | ✓ 1070ms | ✓ 1436ms | ✓ 1225ms | http |
| 101.43.255.96:80 | ✓ 1027ms | ✓ 1309ms | ✓ 1001ms | 否 | ✓ 940ms | http |
| 45.129.141.143:3128 | ✓ 1100ms | 否 | ✓ 1625ms | ✓ 1902ms | ✓ 1593ms | http |
| 190.9.109.198:999 | ✓ 952ms | 否 | ✓ 1239ms | ✓ 1351ms | ✓ 1472ms | http |
| 120.198.141.79:22222 | ✓ 1047ms | ✓ 1283ms | ✓ 1144ms | ✓ 1320ms | ✓ 981ms | http |
| 222.184.48.242:22222 | ✓ 1015ms | ✓ 980ms | ✓ 1075ms | 否 | ✓ 883ms | http |
| 103.84.95.54:7890 | ✓ 1290ms | 否 | ✓ 1778ms | 否 | ✓ 1438ms | http |
| 171.251.172.78:5104 | 否 | 否 | ✓ 1895ms | ✓ 1646ms | ✓ 1648ms | http |
| 120.240.35.178:22222 | 否 | ✓ 1310ms | ✓ 1073ms | ✓ 1288ms | ✓ 1058ms | http |
| 183.249.5.214:22222 | 否 | ✓ 987ms | ✓ 1023ms | ✓ 1068ms | ✓ 789ms | http |
| 183.249.5.111:22222 | 否 | ✓ 1266ms | ✓ 771ms | ✓ 1025ms | ✓ 804ms | http |
| 120.240.35.161:22222 | 否 | ✓ 1324ms | ✓ 1011ms | ✓ 1204ms | ✓ 955ms | http |
| 120.238.159.189:22222 | 否 | ✓ 1312ms | ✓ 997ms | ✓ 1194ms | ✓ 1010ms | http |
| 120.240.35.176:22222 | 否 | ✓ 1275ms | ✓ 985ms | ✓ 1270ms | ✓ 998ms | http |
| 120.240.35.160:22222 | 否 | ✓ 1319ms | ✓ 984ms | ✓ 1215ms | ✓ 998ms | http |
| 113.59.32.141:22222 | 否 | ✓ 1963ms | ✓ 1016ms | ✓ 1231ms | ✓ 1023ms | http |
| 83.219.250.8:62920 | ✓ 882ms | 否 | ✓ 1375ms | 否 | ✓ 1747ms | http |
| 165.227.5.10:8888 | ✓ 1227ms | 否 | ✓ 1551ms | ✓ 1197ms | 否 | http |
| 45.136.130.175:8443 | ✓ 221ms | ✓ 1396ms | ✓ 399ms | 否 | 否 | http |
| 35.225.22.61:80 | ✓ 378ms | 否 | ✓ 1923ms | ✓ 942ms | 否 | http |
| 168.235.110.63:3128 | ✓ 826ms | ✓ 1580ms | ✓ 1656ms | ✓ 1836ms | ✓ 1203ms | http |
| 62.113.119.14:8080 | ✓ 739ms | 否 | ✓ 665ms | ✓ 1549ms | ✓ 1152ms | http |
| 46.183.25.8:443 | ✓ 741ms | 否 | ✓ 658ms | ✓ 1213ms | 否 | http |
| 45.136.130.223:8443 | ✓ 254ms | 否 | ✓ 193ms | ✓ 846ms | ✓ 842ms | http |
| 45.140.147.82:1081 | ✓ 910ms | ✓ 1560ms | ✓ 971ms | ✓ 1605ms | ✓ 1377ms | http |
| 45.136.198.40:3128 | ✓ 1174ms | ✓ 1921ms | 否 | 否 | ✓ 1876ms | http |
| 45.136.130.188:8443 | ✓ 202ms | ✓ 1140ms | ✓ 210ms | ✓ 835ms | ✓ 647ms | http |
| 45.136.130.191:8443 | ✓ 200ms | ✓ 1111ms | ✓ 196ms | 否 | ✓ 630ms | http |
| 185.191.236.162:3128 | ✓ 642ms | 否 | ✓ 606ms | ✓ 1671ms | ✓ 1182ms | http |
| 210.223.44.230:3128 | 否 | ✓ 1066ms | ✓ 869ms | ✓ 1404ms | ✓ 801ms | http |
| 138.124.53.25:7443 | ✓ 791ms | 否 | ✓ 1340ms | 否 | ✓ 1770ms | http |
| 120.238.159.230:22222 | ✓ 1035ms | ✓ 1369ms | ✓ 1154ms | ✓ 1240ms | ✓ 1012ms | http |
| 152.42.213.210:443 | ✓ 1571ms | 否 | ✓ 1272ms | ✓ 1702ms | 否 | http |
| 47.77.193.180:1080 | 否 | ✓ 1563ms | ✓ 235ms | ✓ 1083ms | ✓ 765ms | http |
| 221.122.91.36:11195 | ✓ 866ms | ✓ 1096ms | ✓ 844ms | ✓ 1200ms | ✓ 940ms | http |
| 160.238.65.4:3128 | ✓ 572ms | 否 | ✓ 611ms | 否 | ✓ 1899ms | http |
| 160.238.65.5:3128 | ✓ 1226ms | 否 | ✓ 936ms | 否 | ✓ 1942ms | http |
| 160.238.65.9:3128 | ✓ 675ms | 否 | ✓ 579ms | ✓ 1582ms | ✓ 1111ms | http |
| 120.198.141.75:22222 | ✓ 1020ms | ✓ 1410ms | ✓ 1112ms | 否 | ✓ 1064ms | http |
| 194.5.212.40:8080 | ✓ 535ms | ✓ 1594ms | ✓ 1318ms | ✓ 1677ms | ✓ 1344ms | http |
| 113.59.32.162:22222 | ✓ 1138ms | ✓ 1274ms | ✓ 1005ms | ✓ 1286ms | ✓ 946ms | http |
| 183.249.5.109:22222 | ✓ 857ms | ✓ 981ms | ✓ 896ms | ✓ 1028ms | ✓ 885ms | http |
| 120.240.29.51:22222 | 否 | ✓ 1608ms | ✓ 958ms | ✓ 1251ms | ✓ 1011ms | http |
| 180.127.149.244:1080 | ✓ 1012ms | ✓ 1038ms | ✓ 964ms | 否 | 否 | http |
| 183.249.5.117:22222 | ✓ 903ms | ✓ 1272ms | ✓ 1000ms | ✓ 1076ms | ✓ 976ms | http |
| 117.159.239.51:22222 | ✓ 994ms | ✓ 1167ms | ✓ 898ms | 否 | 否 | http |
| 160.238.65.3:3128 | ✓ 1661ms | 否 | ✓ 1218ms | ✓ 1432ms | ✓ 1105ms | http |
| 160.238.65.8:3128 | ✓ 1661ms | 否 | ✓ 1225ms | ✓ 1433ms | ✓ 1121ms | http |
| 160.238.65.7:3128 | ✓ 1661ms | 否 | ✓ 1219ms | ✓ 1442ms | ✓ 1122ms | http |
| 45.136.131.47:8443 | ✓ 1274ms | ✓ 762ms | ✓ 240ms | ✓ 984ms | 否 | http |
| 171.251.172.78:5102 | ✓ 1974ms | 否 | 否 | ✓ 1667ms | ✓ 1687ms | http |
| 183.249.5.110:22222 | 否 | ✓ 1345ms | ✓ 1997ms | 否 | ✓ 842ms | http |
| 160.250.4.245:1 | 否 | 否 | ✓ 1518ms | ✓ 1687ms | ✓ 1055ms | http |
| 221.122.91.36:11273 | ✓ 947ms | ✓ 1090ms | ✓ 972ms | ✓ 1959ms | ✓ 927ms | http |
| 34.101.184.164:3128 | ✓ 1627ms | 否 | ✓ 955ms | ✓ 1307ms | ✓ 1042ms | http |
| 59.46.216.131:30001 | 否 | ✓ 1348ms | ✓ 1356ms | ✓ 1344ms | 否 | http |
| 222.184.48.248:22222 | ✓ 1064ms | 否 | ✓ 1797ms | ✓ 1476ms | 否 | http |
| 45.186.6.104:3128 | ✓ 1899ms | ✓ 1663ms | ✓ 1832ms | 否 | 否 | http |
| 120.238.159.229:22222 | ✓ 1671ms | 否 | ✓ 1259ms | ✓ 1295ms | ✓ 986ms | http |
| 46.39.105.157:8080 | ✓ 1880ms | 否 | ✓ 1925ms | ✓ 1483ms | ✓ 1264ms | http |
| 178.236.245.59:3128 | ✓ 1435ms | 否 | 否 | ✓ 1969ms | ✓ 1577ms | http |
| 178.236.245.17:3128 | ✓ 1463ms | 否 | 否 | ✓ 1949ms | ✓ 1596ms | http |
| 222.184.48.251:22222 | 否 | ✓ 1142ms | ✓ 1046ms | 否 | ✓ 891ms | http |
| 16.78.119.130:443 | 否 | ✓ 1944ms | 否 | ✓ 1987ms | ✓ 1674ms | http |
| 120.198.141.80:22222 | 否 | ✓ 1363ms | ✓ 1073ms | 否 | ✓ 1022ms | http |
| 120.240.29.174:22222 | ✓ 1015ms | 否 | ✓ 1072ms | 否 | ✓ 1060ms | http |
| 61.52.131.172:8443 | ✓ 820ms | ✓ 1407ms | ✓ 853ms | 否 | ✓ 856ms | http |
| 45.140.147.155:1082 | ✓ 1041ms | 否 | ✓ 996ms | ✓ 1886ms | ✓ 1252ms | http |
| 171.251.172.78:5107 | ✓ 1598ms | 否 | 否 | ✓ 1799ms | ✓ 1420ms | http |
| 106.117.208.101:7890 | ✓ 982ms | ✓ 1175ms | 否 | ✓ 1450ms | ✓ 1002ms | http |
| 45.140.147.155:1081 | ✓ 701ms | 否 | ✓ 1502ms | ✓ 1737ms | ✓ 1220ms | http |

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
