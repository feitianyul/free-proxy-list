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

最后更新：2026-05-17 00:46:54 UTC（2026-05-17 08:46:54 UTC+8）

**代理总数：72**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 72 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 72 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 113.160.132.26:8080 | 否 | ✓ 1414ms | ✓ 911ms | ✓ 1281ms | ✓ 1014ms | http |
| 212.58.132.5:8888 | ✓ 1209ms | 否 | ✓ 1150ms | ✓ 1561ms | ✓ 1196ms | http |
| 137.59.47.73:3128 | ✓ 1061ms | ✓ 1528ms | ✓ 1330ms | ✓ 1304ms | ✓ 1113ms | http |
| 45.125.67.37:8443 | ✓ 892ms | 否 | ✓ 1409ms | 否 | ✓ 1550ms | http |
| 170.106.136.181:31002 | ✓ 1182ms | ✓ 768ms | ✓ 1121ms | 否 | 否 | http |
| 59.46.216.131:30001 | ✓ 1076ms | ✓ 1389ms | ✓ 1051ms | ✓ 1297ms | 否 | http |
| 91.233.223.147:3128 | ✓ 971ms | ✓ 1996ms | ✓ 884ms | 否 | ✓ 1591ms | http |
| 185.200.188.234:10001 | ✓ 1271ms | 否 | ✓ 976ms | 否 | ✓ 1630ms | http |
| 190.93.224.32:999 | ✓ 1514ms | 否 | ✓ 1128ms | 否 | ✓ 1755ms | http |
| 210.223.44.230:3128 | ✓ 1466ms | 否 | ✓ 1759ms | ✓ 1033ms | 否 | http |
| 120.92.212.16:8890 | ✓ 967ms | ✓ 1150ms | ✓ 998ms | 否 | 否 | http |
| 8.154.21.175:3128 | ✓ 805ms | ✓ 1038ms | ✓ 889ms | ✓ 1114ms | ✓ 941ms | http |
| 103.21.220.141:3128 | ✓ 658ms | 否 | ✓ 672ms | ✓ 1404ms | ✓ 673ms | http |
| 148.230.4.241:999 | ✓ 965ms | ✓ 1501ms | ✓ 478ms | ✓ 1395ms | ✓ 1132ms | http |
| 128.199.114.189:9090 | ✓ 942ms | 否 | ✓ 1354ms | ✓ 1345ms | ✓ 1084ms | http |
| 45.129.141.143:3128 | ✓ 841ms | 否 | ✓ 1939ms | 否 | ✓ 1782ms | http |
| 185.71.196.92:1080 | ✓ 978ms | 否 | ✓ 1735ms | 否 | ✓ 1874ms | http |
| 34.96.238.40:8080 | 否 | ✓ 1213ms | ✓ 1830ms | 否 | ✓ 976ms | http |
| 129.213.162.27:17777 | ✓ 1707ms | ✓ 1353ms | ✓ 1909ms | 否 | 否 | http |
| 51.161.50.166:3128 | ✓ 713ms | ✓ 1706ms | ✓ 1727ms | 否 | 否 | http |
| 84.47.150.125:1080 | ✓ 1113ms | ✓ 1989ms | 否 | 否 | ✓ 1714ms | http |
| 77.110.107.80:1080 | ✓ 1100ms | 否 | ✓ 1562ms | ✓ 1774ms | 否 | http |
| 1.231.81.166:3128 | 否 | ✓ 1443ms | ✓ 816ms | ✓ 1033ms | ✓ 835ms | http |
| 115.231.181.40:8128 | ✓ 1291ms | 否 | 否 | ✓ 1163ms | ✓ 1136ms | http |
| 34.101.184.164:3128 | 否 | 否 | ✓ 1779ms | ✓ 1946ms | ✓ 1554ms | http |
| 180.103.19.252:1080 | ✓ 1203ms | ✓ 1944ms | ✓ 1130ms | 否 | ✓ 1234ms | http |
| 150.107.140.238:3128 | ✓ 1441ms | 否 | ✓ 860ms | ✓ 1111ms | 否 | http |
| 120.92.212.16:7890 | ✓ 885ms | 否 | ✓ 1103ms | 否 | ✓ 1001ms | http |
| 114.214.170.41:27890 | 否 | 否 | ✓ 1244ms | ✓ 1356ms | ✓ 1113ms | http |
| 91.242.229.129:8092 | ✓ 1934ms | 否 | ✓ 1141ms | 否 | ✓ 1937ms | http |
| 114.214.165.78:10810 | ✓ 1415ms | ✓ 1685ms | 否 | ✓ 1385ms | ✓ 1526ms | http |
| 168.138.202.218:3128 | ✓ 1629ms | ✓ 1201ms | ✓ 1200ms | ✓ 1275ms | ✓ 1027ms | http |
| 146.56.164.121:3128 | ✓ 1671ms | ✓ 1929ms | ✓ 1041ms | ✓ 1166ms | ✓ 771ms | http |
| 121.230.9.212:1080 | ✓ 1035ms | ✓ 1379ms | ✓ 1196ms | ✓ 1375ms | ✓ 1072ms | http |
| 121.230.8.72:1080 | ✓ 1174ms | ✓ 1313ms | ✓ 1115ms | ✓ 1421ms | ✓ 1321ms | http |
| 103.235.67.190:80 | ✓ 1203ms | 否 | ✓ 1727ms | ✓ 1567ms | ✓ 973ms | http |
| 129.80.217.21:444 | ✓ 1042ms | ✓ 1154ms | ✓ 1118ms | 否 | ✓ 1933ms | http |
| 210.76.192.50:10808 | 否 | 否 | ✓ 1999ms | ✓ 1562ms | ✓ 1537ms | http |
| 129.80.238.83:444 | ✓ 1038ms | ✓ 1156ms | ✓ 1120ms | ✓ 1119ms | 否 | http |
| 202.40.186.222:27202 | 否 | 否 | ✓ 1866ms | ✓ 1722ms | ✓ 1635ms | http |
| 152.32.132.190:7890 | ✓ 1967ms | 否 | 否 | ✓ 795ms | ✓ 1651ms | http |
| 121.147.253.205:3036 | ✓ 1753ms | ✓ 1058ms | ✓ 765ms | ✓ 1022ms | ✓ 1051ms | http |
| 159.223.41.216:9090 | ✓ 740ms | 否 | ✓ 1248ms | ✓ 1056ms | ✓ 841ms | http |
| 101.32.243.189:80 | ✓ 1209ms | 否 | ✓ 1235ms | 否 | ✓ 1345ms | http |
| 159.89.31.62:8080 | ✓ 651ms | 否 | ✓ 1836ms | 否 | ✓ 1791ms | http |
| 3.101.133.120:80 | ✓ 429ms | ✓ 1088ms | ✓ 154ms | ✓ 991ms | ✓ 976ms | http |
| 14.242.161.183:5106 | ✓ 1356ms | ✓ 1944ms | ✓ 1003ms | ✓ 1995ms | ✓ 1280ms | http |
| 38.210.201.92:999 | 否 | ✓ 1640ms | ✓ 1972ms | 否 | ✓ 1166ms | http |
| 160.238.65.5:3128 | ✓ 1605ms | ✓ 1967ms | ✓ 634ms | ✓ 1627ms | 否 | http |
| 139.162.153.201:3128 | ✓ 913ms | 否 | ✓ 1735ms | 否 | ✓ 1947ms | http |
| 103.180.126.146:8080 | ✓ 1326ms | 否 | ✓ 1888ms | ✓ 1394ms | ✓ 1399ms | http |
| 64.188.77.221:3128 | 否 | ✓ 1584ms | ✓ 1071ms | ✓ 1826ms | ✓ 1499ms | http |
| 43.156.90.221:10808 | 否 | 否 | ✓ 1266ms | ✓ 1020ms | ✓ 792ms | http |
| 178.63.155.151:8888 | ✓ 833ms | ✓ 1966ms | ✓ 829ms | 否 | 否 | http |
| 168.110.52.228:3128 | ✓ 1572ms | 否 | ✓ 1153ms | ✓ 963ms | ✓ 679ms | http |
| 64.181.240.152:3128 | 否 | ✓ 1827ms | ✓ 261ms | ✓ 842ms | 否 | http |
| 157.0.142.246:10057 | ✓ 1102ms | ✓ 1270ms | ✓ 1080ms | ✓ 1375ms | 否 | http |
| 185.21.15.206:3128 | ✓ 1176ms | 否 | 否 | ✓ 1627ms | ✓ 1849ms | http |
| 107.175.85.198:1080 | ✓ 1689ms | ✓ 1363ms | ✓ 865ms | ✓ 1383ms | ✓ 1350ms | http |
| 8.219.97.248:80 | ✓ 1430ms | 否 | ✓ 1384ms | ✓ 1289ms | 否 | http |
| 158.160.215.167:8124 | ✓ 1627ms | 否 | ✓ 1387ms | 否 | ✓ 1700ms | http |
| 103.147.152.12:1080 | ✓ 1634ms | ✓ 1684ms | ✓ 1922ms | 否 | 否 | http |
| 45.88.0.114:3128 | ✓ 1282ms | 否 | ✓ 612ms | 否 | ✓ 1173ms | http |
| 128.199.113.85:9090 | ✓ 1077ms | 否 | 否 | ✓ 1932ms | ✓ 974ms | http |
| 128.199.116.219:9090 | ✓ 1974ms | 否 | ✓ 1554ms | 否 | ✓ 1944ms | http |
| 103.147.152.12:1095 | ✓ 1042ms | ✓ 1715ms | ✓ 865ms | 否 | 否 | http |
| 61.52.131.172:8443 | ✓ 870ms | 否 | ✓ 878ms | ✓ 1247ms | ✓ 1018ms | http |
| 104.247.51.76:3128 | ✓ 1367ms | 否 | ✓ 1257ms | ✓ 1359ms | 否 | http |
| 103.172.70.173:8080 | 否 | 否 | ✓ 1760ms | ✓ 1469ms | ✓ 1363ms | http |
| 38.210.201.96:999 | 否 | ✓ 1347ms | ✓ 1982ms | 否 | ✓ 1384ms | http |
| 103.157.117.116:8080 | 否 | 否 | ✓ 1591ms | ✓ 1748ms | ✓ 1668ms | http |
| 138.2.239.213:10010 | ✓ 1511ms | ✓ 1598ms | ✓ 1412ms | ✓ 1487ms | ✓ 700ms | http |

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
