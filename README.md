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

最后更新：2026-03-14 13:57:28 UTC（2026-03-14 21:57:28 UTC+8）

**代理总数：68**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 68 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 68 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 202.155.12.161:443 | ✓ 1785ms | 否 | ✓ 969ms | ✓ 1390ms | 否 | http |
| 205.209.118.30:3138 | ✓ 266ms | 否 | ✓ 1084ms | ✓ 1130ms | ✓ 863ms | http |
| 150.230.249.50:1080 | ✓ 1819ms | 否 | ✓ 1992ms | ✓ 1271ms | ✓ 1939ms | http |
| 47.77.193.180:1080 | ✓ 962ms | ✓ 1462ms | ✓ 330ms | ✓ 918ms | ✓ 659ms | http |
| 38.180.2.107:3128 | ✓ 874ms | ✓ 1994ms | ✓ 1806ms | 否 | 否 | http |
| 223.16.170.103:80 | ✓ 1368ms | 否 | ✓ 1529ms | ✓ 1556ms | 否 | http |
| 35.225.22.61:80 | ✓ 1151ms | 否 | ✓ 1019ms | 否 | ✓ 785ms | http |
| 101.43.127.100:8877 | ✓ 783ms | 否 | ✓ 1443ms | ✓ 1272ms | 否 | http |
| 210.77.29.245:7890 | ✓ 1722ms | 否 | ✓ 1371ms | ✓ 1485ms | ✓ 1102ms | http |
| 85.198.96.242:3128 | ✓ 581ms | ✓ 1849ms | ✓ 1229ms | ✓ 1626ms | ✓ 1323ms | http |
| 45.136.198.40:3128 | ✓ 1485ms | 否 | ✓ 1792ms | 否 | ✓ 1653ms | http |
| 81.70.169.194:80 | ✓ 1385ms | ✓ 1272ms | 否 | ✓ 1510ms | 否 | http |
| 113.160.132.26:8080 | ✓ 1534ms | ✓ 1464ms | ✓ 1159ms | ✓ 1385ms | ✓ 1084ms | http |
| 20.210.76.175:8561 | ✓ 613ms | ✓ 1093ms | ✓ 653ms | ✓ 1082ms | ✓ 1457ms | http |
| 20.210.76.104:8561 | ✓ 613ms | ✓ 1101ms | ✓ 691ms | ✓ 1038ms | ✓ 1457ms | http |
| 20.210.76.178:8561 | ✓ 609ms | ✓ 1719ms | ✓ 644ms | ✓ 1279ms | ✓ 1868ms | http |
| 20.78.26.206:8561 | ✓ 1437ms | ✓ 1339ms | ✓ 870ms | ✓ 978ms | ✓ 797ms | http |
| 20.210.39.153:8561 | ✓ 1433ms | ✓ 1400ms | ✓ 809ms | ✓ 977ms | ✓ 812ms | http |
| 20.78.118.91:8561 | ✓ 1437ms | 否 | ✓ 664ms | ✓ 967ms | ✓ 768ms | http |
| 103.84.95.54:7890 | ✓ 1018ms | 否 | ✓ 802ms | 否 | ✓ 773ms | http |
| 62.60.177.204:34094 | ✓ 709ms | ✓ 1969ms | ✓ 1664ms | 否 | 否 | http |
| 34.96.238.40:8080 | 否 | 否 | ✓ 1263ms | ✓ 1143ms | ✓ 1429ms | http |
| 20.27.13.35:8561 | ✓ 1501ms | ✓ 1971ms | ✓ 644ms | ✓ 996ms | ✓ 738ms | http |
| 20.27.14.220:8561 | ✓ 1549ms | ✓ 1069ms | ✓ 595ms | ✓ 967ms | ✓ 757ms | http |
| 120.92.212.16:8890 | ✓ 1294ms | ✓ 1931ms | 否 | ✓ 1251ms | 否 | http |
| 43.167.227.161:1080 | 否 | 否 | ✓ 601ms | ✓ 963ms | ✓ 757ms | http |
| 120.92.212.16:7890 | 否 | 否 | ✓ 1065ms | ✓ 1271ms | ✓ 1011ms | http |
| 223.205.212.157:8080 | ✓ 1876ms | 否 | ✓ 1846ms | ✓ 1785ms | 否 | http |
| 104.243.46.122:3128 | ✓ 1067ms | ✓ 1356ms | ✓ 1583ms | ✓ 1057ms | 否 | http |
| 121.40.231.103:7890 | 否 | ✓ 1008ms | ✓ 851ms | ✓ 1049ms | ✓ 784ms | http |
| 91.233.223.147:3128 | ✓ 967ms | 否 | ✓ 778ms | ✓ 1939ms | ✓ 1512ms | http |
| 45.149.92.147:5001 | ✓ 1481ms | 否 | ✓ 786ms | ✓ 941ms | ✓ 763ms | http |
| 113.176.92.71:3128 | ✓ 1569ms | ✓ 1505ms | ✓ 1417ms | 否 | ✓ 1076ms | http |
| 106.117.208.101:7890 | ✓ 1032ms | ✓ 1328ms | ✓ 1021ms | ✓ 1347ms | 否 | http |
| 34.101.184.164:3128 | ✓ 1918ms | 否 | ✓ 1466ms | ✓ 1980ms | ✓ 1180ms | http |
| 195.86.215.2:3128 | ✓ 960ms | 否 | ✓ 956ms | ✓ 1208ms | ✓ 966ms | http |
| 210.223.44.230:3128 | 否 | 否 | ✓ 755ms | ✓ 1086ms | ✓ 887ms | http |
| 120.238.159.234:22222 | ✓ 1128ms | ✓ 1321ms | ✓ 1177ms | ✓ 1272ms | ✓ 1004ms | http |
| 207.254.71.62:8088 | ✓ 683ms | 否 | ✓ 1483ms | 否 | ✓ 1948ms | http |
| 128.199.120.45:9090 | 否 | 否 | ✓ 1166ms | ✓ 1548ms | ✓ 1239ms | http |
| 38.145.218.102:8443 | ✓ 955ms | ✓ 950ms | ✓ 474ms | ✓ 1813ms | ✓ 961ms | http |
| 20.27.11.248:8561 | ✓ 725ms | ✓ 1316ms | ✓ 723ms | ✓ 963ms | ✓ 731ms | http |
| 20.27.15.111:8561 | ✓ 726ms | ✓ 1268ms | ✓ 771ms | ✓ 964ms | ✓ 730ms | http |
| 45.136.130.217:8443 | ✓ 949ms | ✓ 1099ms | ✓ 1048ms | ✓ 1232ms | ✓ 908ms | http |
| 45.136.130.215:8443 | ✓ 937ms | ✓ 1649ms | ✓ 505ms | ✓ 1230ms | ✓ 886ms | http |
| 45.136.130.245:8447 | ✓ 956ms | ✓ 1042ms | ✓ 1257ms | ✓ 1200ms | ✓ 932ms | http |
| 45.136.130.214:8443 | ✓ 933ms | 否 | ✓ 496ms | ✓ 1451ms | ✓ 912ms | http |
| 45.136.130.216:8443 | ✓ 1217ms | 否 | ✓ 844ms | ✓ 1194ms | ✓ 872ms | http |
| 45.136.131.39:8443 | 否 | 否 | ✓ 488ms | ✓ 1201ms | ✓ 1055ms | http |
| 47.101.159.19:8899 | ✓ 912ms | ✓ 995ms | ✓ 1051ms | 否 | ✓ 858ms | http |
| 163.44.126.97:3128 | ✓ 1401ms | 否 | 否 | ✓ 1877ms | ✓ 1131ms | http |
| 138.124.53.25:7443 | ✓ 1982ms | 否 | ✓ 1492ms | 否 | ✓ 1486ms | http |
| 101.43.255.96:80 | 否 | ✓ 1523ms | ✓ 1097ms | ✓ 1887ms | ✓ 1018ms | http |
| 202.38.72.235:26001 | ✓ 1871ms | 否 | ✓ 1797ms | ✓ 1886ms | 否 | http |
| 223.16.170.103:3128 | ✓ 1275ms | 否 | ✓ 1203ms | ✓ 1237ms | ✓ 1274ms | http |
| 103.113.70.189:1081 | ✓ 283ms | 否 | 否 | ✓ 1296ms | ✓ 1032ms | http |
| 14.225.212.37:7890 | ✓ 1657ms | ✓ 1562ms | ✓ 971ms | ✓ 1249ms | ✓ 975ms | http |
| 38.145.218.82:8443 | ✓ 936ms | ✓ 1168ms | ✓ 1009ms | ✓ 964ms | ✓ 758ms | http |
| 168.235.110.63:3128 | ✓ 622ms | 否 | ✓ 1133ms | ✓ 1064ms | ✓ 778ms | http |
| 38.145.203.235:8443 | ✓ 920ms | 否 | ✓ 618ms | ✓ 1073ms | ✓ 717ms | http |
| 45.136.130.232:8443 | ✓ 904ms | ✓ 1573ms | ✓ 1033ms | ✓ 1092ms | ✓ 735ms | http |
| 38.145.208.133:8443 | ✓ 902ms | 否 | ✓ 619ms | ✓ 1100ms | ✓ 729ms | http |
| 38.145.208.137:8443 | ✓ 471ms | ✓ 1735ms | ✓ 474ms | ✓ 1123ms | ✓ 737ms | http |
| 101.47.73.135:3128 | ✓ 1224ms | 否 | ✓ 1122ms | 否 | ✓ 1337ms | http |
| 38.145.208.131:8443 | ✓ 563ms | ✓ 1295ms | ✓ 1040ms | ✓ 973ms | ✓ 734ms | http |
| 47.101.149.27:9010 | ✓ 1319ms | ✓ 1195ms | 否 | 否 | ✓ 1191ms | http |
| 59.46.216.131:30001 | ✓ 1075ms | ✓ 1404ms | 否 | 否 | ✓ 1865ms | http |
| 159.223.42.219:3128 | ✓ 1610ms | 否 | ✓ 1455ms | ✓ 1473ms | ✓ 1071ms | http |

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
