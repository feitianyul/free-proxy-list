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

最后更新：2026-03-14 16:27:42 UTC（2026-03-15 00:27:42 UTC+8）

**代理总数：88**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 88 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 88 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 989ms | 否 | ✓ 1084ms | ✓ 1129ms | ✓ 857ms | http |
| 113.160.132.26:8080 | ✓ 1959ms | 否 | ✓ 1582ms | ✓ 1419ms | ✓ 1454ms | http |
| 103.84.95.54:7890 | ✓ 847ms | 否 | ✓ 1325ms | ✓ 1916ms | ✓ 1353ms | http |
| 150.230.249.50:1080 | ✓ 1917ms | 否 | ✓ 967ms | 否 | ✓ 1928ms | http |
| 45.149.92.147:5001 | ✓ 1774ms | 否 | ✓ 795ms | ✓ 967ms | ✓ 1588ms | http |
| 113.59.32.160:22222 | ✓ 1150ms | ✓ 1460ms | ✓ 1137ms | ✓ 1424ms | ✓ 1068ms | http |
| 120.92.212.16:7890 | 否 | ✓ 1419ms | ✓ 1077ms | ✓ 1437ms | 否 | http |
| 38.145.218.82:8443 | ✓ 1321ms | ✓ 1221ms | ✓ 1173ms | ✓ 959ms | 否 | http |
| 38.145.203.235:8443 | ✓ 714ms | 否 | ✓ 951ms | ✓ 945ms | ✓ 759ms | http |
| 38.55.107.254:6005 | 否 | 否 | ✓ 1101ms | ✓ 1240ms | ✓ 998ms | http |
| 101.43.255.96:80 | ✓ 1238ms | ✓ 1481ms | ✓ 1169ms | 否 | 否 | http |
| 38.145.203.135:8443 | 否 | ✓ 1159ms | ✓ 1378ms | ✓ 1111ms | ✓ 903ms | http |
| 86.53.183.16:1080 | ✓ 439ms | 否 | ✓ 1886ms | 否 | ✓ 1438ms | http |
| 62.60.177.204:34094 | ✓ 914ms | 否 | 否 | ✓ 1028ms | ✓ 889ms | http |
| 35.225.22.61:80 | ✓ 621ms | 否 | ✓ 130ms | 否 | ✓ 1062ms | http |
| 45.136.131.39:8443 | ✓ 609ms | 否 | ✓ 332ms | ✓ 945ms | ✓ 748ms | http |
| 120.92.212.16:8890 | ✓ 1080ms | 否 | ✓ 1094ms | ✓ 1683ms | 否 | http |
| 45.136.131.42:8447 | ✓ 317ms | ✓ 1444ms | ✓ 440ms | ✓ 968ms | ✓ 735ms | http |
| 38.145.218.102:8443 | ✓ 314ms | ✓ 1812ms | ✓ 316ms | ✓ 938ms | ✓ 734ms | http |
| 183.249.5.109:22222 | ✓ 834ms | 否 | ✓ 882ms | ✓ 1094ms | ✓ 1008ms | http |
| 183.249.5.214:22222 | ✓ 903ms | ✓ 1095ms | ✓ 974ms | ✓ 1132ms | 否 | http |
| 113.59.32.161:22222 | ✓ 1169ms | 否 | ✓ 1009ms | ✓ 1411ms | 否 | http |
| 120.240.35.176:22222 | ✓ 1094ms | 否 | ✓ 1135ms | 否 | ✓ 1084ms | http |
| 43.167.227.161:1080 | 否 | ✓ 1579ms | ✓ 666ms | ✓ 925ms | ✓ 740ms | http |
| 120.240.35.178:22222 | 否 | ✓ 1408ms | ✓ 1185ms | ✓ 1304ms | ✓ 1061ms | http |
| 120.240.35.173:22222 | ✓ 1142ms | ✓ 1400ms | ✓ 1038ms | ✓ 1307ms | ✓ 1063ms | http |
| 150.249.255.91:3128 | ✓ 1705ms | 否 | 否 | ✓ 1546ms | ✓ 812ms | http |
| 81.70.169.194:80 | ✓ 1741ms | 否 | ✓ 1062ms | 否 | ✓ 1436ms | http |
| 101.47.73.135:3128 | ✓ 1729ms | 否 | ✓ 1309ms | ✓ 1802ms | ✓ 1617ms | http |
| 38.145.203.106:8443 | ✓ 331ms | ✓ 1141ms | ✓ 985ms | ✓ 954ms | 否 | http |
| 38.145.203.105:8443 | ✓ 326ms | ✓ 1696ms | ✓ 424ms | ✓ 980ms | 否 | http |
| 38.145.203.110:8443 | ✓ 322ms | 否 | ✓ 309ms | ✓ 940ms | 否 | http |
| 38.145.203.124:8443 | ✓ 316ms | 否 | ✓ 331ms | ✓ 951ms | 否 | http |
| 91.233.223.147:3128 | ✓ 838ms | 否 | ✓ 1065ms | ✓ 1927ms | 否 | http |
| 38.145.218.101:8447 | ✓ 327ms | ✓ 1833ms | ✓ 1782ms | 否 | 否 | http |
| 162.240.154.26:3128 | ✓ 1584ms | 否 | ✓ 1310ms | 否 | ✓ 1804ms | http |
| 168.235.110.63:3128 | ✓ 1375ms | 否 | ✓ 709ms | 否 | ✓ 1136ms | http |
| 183.249.5.117:22222 | ✓ 1143ms | ✓ 1036ms | ✓ 1069ms | ✓ 1078ms | ✓ 862ms | http |
| 120.240.35.160:22222 | ✓ 1067ms | ✓ 1417ms | ✓ 1193ms | ✓ 1360ms | ✓ 1098ms | http |
| 210.223.44.230:3128 | ✓ 1295ms | ✓ 977ms | ✓ 722ms | ✓ 1058ms | ✓ 820ms | http |
| 45.140.147.82:1081 | ✓ 415ms | ✓ 1940ms | ✓ 776ms | ✓ 1517ms | ✓ 1157ms | http |
| 45.140.147.82:1082 | ✓ 482ms | ✓ 1434ms | ✓ 1215ms | ✓ 1513ms | ✓ 1172ms | http |
| 38.145.208.201:8447 | ✓ 377ms | ✓ 913ms | ✓ 315ms | ✓ 965ms | ✓ 757ms | http |
| 45.136.130.245:8447 | 否 | 否 | ✓ 503ms | ✓ 950ms | ✓ 762ms | http |
| 120.240.29.51:22222 | 否 | ✓ 1413ms | ✓ 1165ms | 否 | ✓ 1046ms | http |
| 192.71.213.85:9812 | ✓ 1457ms | 否 | ✓ 1764ms | ✓ 1822ms | 否 | http |
| 120.238.159.229:22222 | ✓ 1136ms | 否 | ✓ 1059ms | ✓ 1269ms | 否 | http |
| 113.176.92.71:3128 | 否 | 否 | ✓ 1478ms | ✓ 1681ms | ✓ 1082ms | http |
| 120.238.159.228:22222 | ✓ 1088ms | 否 | 否 | ✓ 1363ms | ✓ 1047ms | http |
| 62.113.119.14:8080 | ✓ 1538ms | 否 | 否 | ✓ 1617ms | ✓ 1231ms | http |
| 106.117.208.101:7890 | 否 | ✓ 1873ms | ✓ 1174ms | ✓ 1371ms | ✓ 1517ms | http |
| 45.136.130.211:8447 | ✓ 967ms | ✓ 1393ms | ✓ 1023ms | ✓ 969ms | ✓ 866ms | http |
| 38.145.218.163:8443 | ✓ 1000ms | 否 | ✓ 656ms | ✓ 1076ms | ✓ 739ms | http |
| 38.145.218.235:8443 | ✓ 998ms | ✓ 1760ms | ✓ 888ms | ✓ 1091ms | ✓ 732ms | http |
| 38.145.218.160:8443 | ✓ 994ms | ✓ 1672ms | ✓ 980ms | ✓ 1102ms | ✓ 742ms | http |
| 38.145.218.134:8443 | ✓ 383ms | ✓ 1372ms | ✓ 390ms | ✓ 968ms | ✓ 765ms | http |
| 38.145.218.161:8443 | ✓ 383ms | ✓ 970ms | ✓ 325ms | ✓ 928ms | ✓ 742ms | http |
| 38.145.208.138:8447 | 否 | 否 | ✓ 911ms | ✓ 990ms | ✓ 829ms | http |
| 120.238.159.250:22222 | ✓ 1145ms | 否 | ✓ 1024ms | ✓ 1286ms | ✓ 1022ms | http |
| 38.145.218.162:8443 | ✓ 367ms | ✓ 960ms | ✓ 622ms | ✓ 944ms | ✓ 723ms | http |
| 222.184.48.251:22222 | ✓ 1026ms | ✓ 1330ms | ✓ 994ms | ✓ 1378ms | 否 | http |
| 14.225.212.37:7890 | ✓ 1686ms | 否 | ✓ 978ms | ✓ 1488ms | ✓ 1114ms | http |
| 116.80.65.81:3172 | ✓ 1641ms | 否 | ✓ 1690ms | 否 | ✓ 1816ms | http |
| 4.247.152.147:3128 | ✓ 1547ms | 否 | ✓ 1445ms | 否 | ✓ 1689ms | http |
| 165.227.5.10:8888 | ✓ 1565ms | 否 | 否 | ✓ 1073ms | ✓ 776ms | http |
| 172.212.68.37:3128 | ✓ 668ms | ✓ 1605ms | ✓ 710ms | ✓ 1867ms | ✓ 912ms | http |
| 223.16.170.103:3128 | ✓ 1634ms | 否 | ✓ 1562ms | ✓ 1256ms | 否 | http |
| 117.159.239.52:22222 | ✓ 939ms | ✓ 1187ms | ✓ 1079ms | ✓ 1357ms | ✓ 996ms | http |
| 222.184.48.252:22222 | ✓ 1912ms | ✓ 1512ms | 否 | 否 | ✓ 1121ms | http |
| 223.16.170.103:80 | ✓ 1427ms | 否 | ✓ 1541ms | 否 | ✓ 1256ms | http |
| 120.232.242.119:22222 | 否 | ✓ 1384ms | 否 | ✓ 1271ms | ✓ 1009ms | http |
| 113.59.32.162:22222 | ✓ 1222ms | ✓ 1444ms | ✓ 1025ms | ✓ 1608ms | 否 | http |
| 24.144.86.173:1080 | ✓ 404ms | ✓ 1646ms | ✓ 825ms | ✓ 870ms | ✓ 766ms | http |
| 85.198.96.242:3128 | ✓ 617ms | 否 | ✓ 1639ms | ✓ 1722ms | 否 | http |
| 45.136.198.40:3128 | ✓ 1031ms | 否 | ✓ 1242ms | 否 | ✓ 1461ms | http |
| 38.145.208.131:8443 | ✓ 767ms | ✓ 971ms | ✓ 779ms | ✓ 963ms | ✓ 792ms | http |
| 84.247.171.137:3128 | ✓ 1002ms | ✓ 1417ms | ✓ 735ms | 否 | ✓ 1163ms | http |
| 38.145.208.185:8443 | 否 | ✓ 1514ms | ✓ 324ms | ✓ 988ms | 否 | http |
| 38.145.218.14:8443 | 否 | ✓ 1511ms | ✓ 329ms | ✓ 996ms | 否 | http |
| 38.145.208.184:8443 | 否 | ✓ 1539ms | ✓ 327ms | ✓ 988ms | 否 | http |
| 38.145.208.181:8443 | 否 | ✓ 1537ms | ✓ 336ms | ✓ 981ms | 否 | http |
| 38.145.208.186:8443 | 否 | ✓ 1548ms | ✓ 338ms | ✓ 970ms | 否 | http |
| 38.145.208.180:8443 | 否 | ✓ 1960ms | ✓ 319ms | ✓ 1734ms | 否 | http |
| 120.238.159.189:22222 | ✓ 1110ms | ✓ 1395ms | ✓ 1125ms | ✓ 1306ms | ✓ 1047ms | http |
| 61.52.131.172:8443 | ✓ 1022ms | ✓ 1350ms | ✓ 1150ms | ✓ 1369ms | ✓ 1017ms | http |
| 210.77.29.245:7890 | ✓ 1017ms | ✓ 1325ms | ✓ 1133ms | ✓ 1320ms | ✓ 1041ms | http |
| 138.124.53.25:7443 | ✓ 1053ms | 否 | ✓ 1992ms | ✓ 1956ms | ✓ 1965ms | http |
| 103.39.51.190:8080 | ✓ 1881ms | 否 | 否 | ✓ 1663ms | ✓ 1822ms | http |

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
