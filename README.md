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

最后更新：2026-03-16 03:42:29 UTC（2026-03-16 11:42:29 UTC+8）

**代理总数：84**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 84 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 84 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 113.160.132.26:8080 | ✓ 1336ms | ✓ 1554ms | ✓ 923ms | ✓ 1538ms | ✓ 983ms | http |
| 45.167.124.52:8080 | ✓ 1017ms | 否 | ✓ 1695ms | 否 | ✓ 1601ms | http |
| 205.209.118.30:3138 | ✓ 1698ms | 否 | ✓ 930ms | ✓ 1252ms | ✓ 932ms | http |
| 133.242.138.34:8100 | ✓ 1475ms | ✓ 1357ms | ✓ 611ms | ✓ 964ms | ✓ 793ms | http |
| 85.198.96.242:3128 | ✓ 1494ms | 否 | ✓ 843ms | ✓ 1894ms | ✓ 1422ms | http |
| 137.220.150.104:6005 | ✓ 1664ms | 否 | ✓ 736ms | ✓ 1319ms | ✓ 1784ms | http |
| 59.46.216.131:30001 | ✓ 1096ms | ✓ 1319ms | 否 | 否 | ✓ 1034ms | http |
| 8.209.239.31:30000 | ✓ 1325ms | ✓ 1781ms | ✓ 449ms | ✓ 717ms | ✓ 557ms | http |
| 47.79.40.38:55000 | ✓ 1322ms | 否 | ✓ 1191ms | ✓ 769ms | ✓ 1848ms | http |
| 137.220.151.110:6005 | 否 | 否 | ✓ 1068ms | ✓ 1282ms | ✓ 1282ms | http |
| 137.220.150.152:6005 | 否 | 否 | ✓ 1929ms | ✓ 1925ms | ✓ 1119ms | http |
| 104.129.202.127:12354 | 否 | 否 | ✓ 539ms | ✓ 1900ms | ✓ 1979ms | http |
| 62.60.177.204:34094 | 否 | ✓ 1336ms | ✓ 1108ms | ✓ 1116ms | 否 | http |
| 120.240.35.176:22222 | ✓ 1543ms | ✓ 1190ms | 否 | ✓ 1232ms | ✓ 1105ms | http |
| 222.184.48.236:22222 | ✓ 987ms | ✓ 1180ms | ✓ 896ms | ✓ 1536ms | 否 | http |
| 101.43.127.100:8877 | ✓ 960ms | ✓ 1704ms | ✓ 1751ms | ✓ 1864ms | ✓ 911ms | http |
| 81.70.169.194:80 | ✓ 975ms | ✓ 1409ms | 否 | ✓ 1449ms | ✓ 1241ms | http |
| 101.43.255.96:80 | ✓ 968ms | 否 | ✓ 1315ms | ✓ 1312ms | ✓ 1011ms | http |
| 168.235.110.63:3128 | ✓ 804ms | 否 | ✓ 869ms | 否 | ✓ 1817ms | http |
| 115.231.181.40:8128 | ✓ 950ms | ✓ 1162ms | 否 | 否 | ✓ 864ms | http |
| 38.145.220.34:8448 | ✓ 685ms | ✓ 657ms | ✓ 84ms | ✓ 664ms | ✓ 580ms | http |
| 38.145.208.162:8448 | ✓ 688ms | ✓ 1520ms | ✓ 122ms | ✓ 872ms | ✓ 1163ms | http |
| 38.145.208.217:8448 | ✓ 710ms | ✓ 616ms | ✓ 1645ms | ✓ 852ms | ✓ 1529ms | http |
| 147.161.210.140:8800 | ✓ 629ms | ✓ 1508ms | ✓ 553ms | ✓ 1192ms | ✓ 1743ms | http |
| 219.117.204.211:7799 | ✓ 849ms | 否 | ✓ 935ms | 否 | ✓ 1775ms | http |
| 149.50.116.240:1080 | ✓ 610ms | 否 | ✓ 1820ms | 否 | ✓ 1533ms | http |
| 38.145.203.135:8443 | ✓ 672ms | ✓ 1054ms | 否 | ✓ 1141ms | ✓ 479ms | http |
| 120.238.159.250:22222 | ✓ 1037ms | ✓ 1190ms | ✓ 1009ms | ✓ 1227ms | ✓ 905ms | http |
| 104.129.202.127:10810 | ✓ 1244ms | 否 | ✓ 1381ms | 否 | ✓ 851ms | http |
| 120.92.212.16:8890 | ✓ 1881ms | ✓ 1422ms | 否 | 否 | ✓ 985ms | http |
| 222.184.48.235:22222 | ✓ 973ms | ✓ 1481ms | ✓ 1558ms | ✓ 1163ms | 否 | http |
| 213.219.214.45:443 | ✓ 817ms | 否 | ✓ 1318ms | 否 | ✓ 1639ms | http |
| 117.159.239.49:22222 | ✓ 765ms | ✓ 1118ms | 否 | ✓ 1068ms | 否 | http |
| 120.92.212.16:7890 | 否 | 否 | ✓ 976ms | ✓ 1680ms | ✓ 1223ms | http |
| 165.227.5.10:8888 | ✓ 1006ms | ✓ 1767ms | 否 | ✓ 1409ms | 否 | http |
| 178.236.245.17:3128 | ✓ 1220ms | 否 | ✓ 866ms | 否 | ✓ 1636ms | http |
| 222.184.48.251:22222 | ✓ 1529ms | ✓ 1200ms | ✓ 910ms | ✓ 1153ms | ✓ 936ms | http |
| 183.249.5.117:22222 | ✓ 1032ms | ✓ 1847ms | ✓ 1099ms | 否 | ✓ 1382ms | http |
| 120.240.35.173:22222 | 否 | 否 | ✓ 1260ms | ✓ 1438ms | ✓ 1524ms | http |
| 35.225.22.61:80 | ✓ 783ms | 否 | ✓ 442ms | 否 | ✓ 1098ms | http |
| 137.220.128.173:8866 | ✓ 1351ms | 否 | ✓ 1783ms | ✓ 1760ms | ✓ 1394ms | http |
| 137.184.1.87:3128 | ✓ 324ms | 否 | ✓ 774ms | ✓ 716ms | ✓ 562ms | http |
| 117.159.239.56:22222 | ✓ 791ms | ✓ 997ms | ✓ 830ms | ✓ 1081ms | ✓ 917ms | http |
| 38.145.208.179:8447 | ✓ 110ms | ✓ 894ms | ✓ 650ms | ✓ 990ms | ✓ 982ms | http |
| 2.56.122.146:10808 | ✓ 1128ms | ✓ 1769ms | ✓ 1041ms | ✓ 1578ms | ✓ 1145ms | http |
| 211.171.114.154:3128 | ✓ 1650ms | ✓ 1379ms | 否 | 否 | ✓ 1839ms | http |
| 137.220.128.149:8866 | ✓ 1636ms | 否 | ✓ 1562ms | ✓ 1728ms | ✓ 1345ms | http |
| 38.180.2.107:3128 | ✓ 1071ms | 否 | ✓ 980ms | ✓ 1815ms | ✓ 1466ms | http |
| 157.0.142.246:10057 | ✓ 1036ms | 否 | ✓ 1231ms | ✓ 1307ms | ✓ 1049ms | http |
| 114.231.73.92:1080 | ✓ 1077ms | 否 | ✓ 1011ms | ✓ 1185ms | ✓ 950ms | http |
| 117.86.6.32:1080 | ✓ 978ms | ✓ 1340ms | ✓ 979ms | 否 | 否 | http |
| 38.145.218.82:8443 | ✓ 460ms | ✓ 865ms | ✓ 469ms | 否 | 否 | http |
| 45.136.198.40:3128 | ✓ 841ms | 否 | ✓ 1722ms | 否 | ✓ 1834ms | http |
| 8.219.97.248:80 | ✓ 931ms | 否 | ✓ 1190ms | ✓ 1175ms | 否 | http |
| 223.16.170.103:80 | ✓ 858ms | 否 | 否 | ✓ 1065ms | ✓ 1391ms | http |
| 103.183.10.172:3125 | ✓ 1741ms | 否 | 否 | ✓ 1467ms | ✓ 1795ms | http |
| 45.88.0.116:3128 | ✓ 600ms | 否 | ✓ 1736ms | ✓ 1569ms | ✓ 1222ms | http |
| 103.84.95.54:7890 | ✓ 661ms | ✓ 1697ms | 否 | ✓ 1128ms | ✓ 766ms | http |
| 117.159.239.46:22222 | ✓ 880ms | ✓ 1171ms | ✓ 842ms | ✓ 1042ms | ✓ 858ms | http |
| 120.238.159.189:22222 | ✓ 1011ms | ✓ 1213ms | ✓ 976ms | ✓ 1162ms | ✓ 901ms | http |
| 45.88.0.115:3128 | 否 | ✓ 1706ms | ✓ 856ms | ✓ 1465ms | ✓ 1137ms | http |
| 120.240.29.51:22222 | 否 | ✓ 1315ms | ✓ 1030ms | 否 | ✓ 901ms | http |
| 117.159.239.51:22222 | ✓ 751ms | 否 | ✓ 815ms | ✓ 1145ms | 否 | http |
| 117.159.239.44:22222 | ✓ 935ms | ✓ 1039ms | ✓ 860ms | ✓ 1078ms | ✓ 828ms | http |
| 101.47.73.135:3128 | ✓ 1084ms | 否 | 否 | ✓ 1124ms | ✓ 1019ms | http |
| 194.5.212.40:8080 | ✓ 717ms | 否 | ✓ 1104ms | 否 | ✓ 1683ms | http |
| 103.155.64.217:8080 | ✓ 1760ms | 否 | ✓ 1694ms | ✓ 1451ms | 否 | http |
| 58.69.201.198:8000 | ✓ 1379ms | 否 | ✓ 1982ms | 否 | ✓ 1881ms | http |
| 124.16.126.149:7893 | ✓ 1076ms | ✓ 1370ms | ✓ 1274ms | ✓ 1367ms | ✓ 1082ms | http |
| 20.210.76.175:8561 | ✓ 1005ms | ✓ 1259ms | ✓ 651ms | ✓ 1041ms | ✓ 597ms | http |
| 117.159.239.54:22222 | 否 | 否 | ✓ 842ms | ✓ 1319ms | ✓ 812ms | http |
| 117.159.239.48:22222 | ✓ 808ms | ✓ 1048ms | ✓ 848ms | ✓ 1051ms | ✓ 865ms | http |
| 120.232.242.120:22222 | ✓ 911ms | ✓ 1224ms | ✓ 954ms | ✓ 1164ms | ✓ 878ms | http |
| 45.136.130.218:8448 | 否 | 否 | ✓ 1797ms | ✓ 1223ms | ✓ 1447ms | http |
| 72.56.79.129:1080 | 否 | ✓ 1961ms | ✓ 1792ms | 否 | ✓ 1351ms | http |
| 103.39.51.190:8080 | 否 | 否 | ✓ 1322ms | ✓ 1458ms | ✓ 1348ms | http |
| 106.117.208.101:7890 | ✓ 983ms | ✓ 1325ms | ✓ 992ms | ✓ 1308ms | ✓ 1020ms | http |
| 45.88.0.114:3128 | ✓ 1909ms | 否 | ✓ 1851ms | ✓ 1965ms | ✓ 1683ms | http |
| 45.88.0.98:3128 | ✓ 1912ms | 否 | ✓ 1848ms | ✓ 1956ms | ✓ 1701ms | http |
| 213.220.62.62:3128 | ✓ 1909ms | ✓ 1941ms | ✓ 1910ms | ✓ 1956ms | ✓ 1693ms | http |
| 45.88.0.113:3128 | ✓ 1909ms | 否 | ✓ 1852ms | ✓ 1978ms | ✓ 1679ms | http |
| 45.88.0.117:3128 | ✓ 1908ms | ✓ 1934ms | ✓ 1917ms | ✓ 1958ms | ✓ 1690ms | http |
| 45.88.0.111:3128 | ✓ 1908ms | 否 | ✓ 1852ms | ✓ 1952ms | ✓ 1695ms | http |
| 45.88.0.99:3128 | ✓ 1291ms | ✓ 1578ms | ✓ 1680ms | ✓ 1953ms | 否 | http |

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
