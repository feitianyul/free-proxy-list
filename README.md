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

最后更新：2026-03-23 18:44:16 UTC（2026-03-24 02:44:16 UTC+8）

**代理总数：78**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 78 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 78 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 8.217.106.71:8888 | ✓ 653ms | ✓ 1232ms | ✓ 666ms | ✓ 845ms | ✓ 673ms | http |
| 147.161.210.140:8800 | ✓ 1469ms | ✓ 1589ms | ✓ 858ms | ✓ 1247ms | ✓ 1173ms | http |
| 113.160.132.26:8080 | ✓ 1497ms | ✓ 1581ms | ✓ 998ms | ✓ 1397ms | ✓ 1098ms | http |
| 167.103.34.108:8800 | ✓ 1445ms | 否 | ✓ 1613ms | ✓ 1603ms | ✓ 1475ms | http |
| 35.225.22.61:80 | ✓ 602ms | ✓ 1339ms | 否 | 否 | ✓ 1146ms | http |
| 43.99.54.236:5555 | ✓ 665ms | ✓ 972ms | ✓ 675ms | ✓ 858ms | ✓ 682ms | http |
| 142.171.224.229:7890 | ✓ 962ms | ✓ 1392ms | ✓ 1730ms | ✓ 1517ms | ✓ 711ms | http |
| 137.220.150.22:6005 | ✓ 832ms | 否 | ✓ 1383ms | ✓ 1153ms | ✓ 917ms | http |
| 167.103.31.122:8800 | ✓ 1772ms | 否 | ✓ 1681ms | 否 | ✓ 1811ms | http |
| 43.225.185.4:8000 | ✓ 976ms | 否 | ✓ 1704ms | ✓ 1317ms | ✓ 1018ms | http |
| 34.101.184.164:3128 | ✓ 1560ms | 否 | ✓ 1201ms | ✓ 1340ms | ✓ 1315ms | http |
| 43.165.195.107:3128 | ✓ 1549ms | ✓ 1795ms | ✓ 971ms | ✓ 1214ms | ✓ 969ms | http |
| 166.88.55.83:7890 | 否 | ✓ 1993ms | ✓ 688ms | ✓ 1005ms | ✓ 797ms | http |
| 147.161.239.240:8800 | ✓ 835ms | ✓ 1822ms | ✓ 931ms | ✓ 1390ms | ✓ 1191ms | http |
| 1.231.81.166:3128 | ✓ 1489ms | ✓ 1121ms | ✓ 1577ms | ✓ 1224ms | ✓ 882ms | http |
| 101.43.127.100:8877 | ✓ 1852ms | 否 | ✓ 1740ms | 否 | ✓ 1670ms | http |
| 101.47.73.135:3128 | 否 | 否 | ✓ 1162ms | ✓ 1721ms | ✓ 1417ms | http |
| 218.60.0.214:80 | ✓ 1061ms | ✓ 1321ms | ✓ 1085ms | ✓ 1273ms | ✓ 1045ms | http |
| 120.92.212.16:8890 | 否 | 否 | ✓ 1046ms | ✓ 1983ms | ✓ 999ms | http |
| 84.247.149.172:3128 | ✓ 805ms | 否 | ✓ 1527ms | ✓ 1525ms | ✓ 966ms | http |
| 103.82.23.118:5253 | ✓ 1432ms | 否 | ✓ 1358ms | ✓ 1516ms | ✓ 1526ms | http |
| 218.89.134.230:3333 | ✓ 1643ms | ✓ 1662ms | ✓ 1636ms | 否 | ✓ 1303ms | http |
| 137.220.150.152:6005 | ✓ 1811ms | 否 | ✓ 1461ms | ✓ 1943ms | 否 | http |
| 120.92.212.16:7890 | ✓ 1929ms | ✓ 1262ms | 否 | ✓ 1281ms | 否 | http |
| 8.212.130.232:8080 | ✓ 814ms | ✓ 1908ms | 否 | 否 | ✓ 901ms | http |
| 115.231.181.40:8128 | ✓ 908ms | ✓ 1136ms | ✓ 911ms | ✓ 1168ms | ✓ 966ms | http |
| 103.139.138.194:3128 | ✓ 1756ms | 否 | ✓ 1178ms | ✓ 1452ms | ✓ 1109ms | http |
| 46.101.190.71:3128 | ✓ 565ms | 否 | ✓ 1304ms | ✓ 1966ms | ✓ 1436ms | http |
| 8.219.97.248:80 | ✓ 1903ms | 否 | ✓ 1557ms | ✓ 1795ms | 否 | http |
| 59.8.203.55:80 | ✓ 1848ms | 否 | ✓ 1160ms | ✓ 995ms | ✓ 788ms | http |
| 219.117.204.211:7799 | 否 | 否 | ✓ 1655ms | ✓ 929ms | ✓ 760ms | http |
| 137.220.150.104:6005 | ✓ 875ms | 否 | ✓ 1034ms | ✓ 1130ms | ✓ 885ms | http |
| 46.39.105.157:8080 | ✓ 892ms | 否 | ✓ 1948ms | 否 | ✓ 1546ms | http |
| 207.254.71.62:8088 | ✓ 1150ms | ✓ 1991ms | 否 | ✓ 1993ms | ✓ 1678ms | http |
| 45.136.198.40:3128 | ✓ 913ms | ✓ 1720ms | 否 | 否 | ✓ 1746ms | http |
| 5.102.109.41:999 | ✓ 1741ms | ✓ 1050ms | ✓ 310ms | 否 | 否 | http |
| 195.123.209.48:3128 | ✓ 1262ms | 否 | ✓ 1827ms | 否 | ✓ 1741ms | http |
| 47.101.159.19:8899 | ✓ 924ms | ✓ 1084ms | ✓ 907ms | ✓ 1144ms | ✓ 891ms | http |
| 166.249.54.61:7234 | ✓ 1469ms | 否 | ✓ 913ms | ✓ 1708ms | ✓ 1919ms | http |
| 181.41.201.85:3128 | ✓ 652ms | 否 | ✓ 1003ms | 否 | ✓ 1781ms | http |
| 45.149.92.147:5001 | ✓ 671ms | 否 | ✓ 645ms | ✓ 829ms | ✓ 664ms | http |
| 65.108.203.35:18080 | ✓ 1614ms | 否 | ✓ 1961ms | 否 | ✓ 1930ms | http |
| 202.141.161.53:30001 | ✓ 1079ms | ✓ 1337ms | ✓ 1174ms | ✓ 1229ms | ✓ 1141ms | http |
| 210.45.70.16:7895 | ✓ 1133ms | ✓ 1526ms | ✓ 1136ms | ✓ 1436ms | ✓ 1137ms | http |
| 106.117.208.101:7890 | ✓ 908ms | ✓ 1266ms | ✓ 1076ms | ✓ 1278ms | ✓ 1053ms | http |
| 209.126.84.189:3128 | ✓ 306ms | ✓ 1269ms | ✓ 974ms | ✓ 1473ms | ✓ 1073ms | http |
| 120.55.163.237:10086 | ✓ 877ms | ✓ 1072ms | ✓ 848ms | ✓ 1083ms | ✓ 874ms | http |
| 160.250.4.245:1 | ✓ 1470ms | 否 | ✓ 1297ms | ✓ 1290ms | ✓ 1035ms | http |
| 192.71.213.85:9812 | ✓ 1166ms | 否 | ✓ 1790ms | ✓ 1899ms | 否 | http |
| 59.46.216.131:30001 | ✓ 995ms | 否 | ✓ 1054ms | ✓ 1376ms | 否 | http |
| 155.212.132.241:3128 | ✓ 1251ms | 否 | 否 | ✓ 1788ms | ✓ 1459ms | http |
| 222.228.171.92:8080 | ✓ 1522ms | 否 | ✓ 1317ms | ✓ 909ms | 否 | http |
| 137.220.151.110:6005 | ✓ 796ms | ✓ 1955ms | ✓ 1009ms | 否 | 否 | http |
| 49.156.44.114:8080 | ✓ 1603ms | 否 | ✓ 1553ms | ✓ 1854ms | ✓ 1517ms | http |
| 178.156.224.42:3128 | ✓ 1649ms | 否 | ✓ 1838ms | 否 | ✓ 1773ms | http |
| 121.230.8.220:1080 | ✓ 1595ms | ✓ 1678ms | ✓ 1340ms | ✓ 1911ms | ✓ 1227ms | http |
| 114.237.77.231:1080 | ✓ 1397ms | 否 | ✓ 1347ms | ✓ 1801ms | ✓ 1466ms | http |
| 103.69.84.106:3131 | ✓ 1931ms | 否 | ✓ 1344ms | ✓ 1582ms | ✓ 1083ms | http |
| 137.184.1.87:3128 | ✓ 725ms | ✓ 1394ms | ✓ 310ms | ✓ 757ms | ✓ 581ms | http |
| 47.238.220.4:8888 | ✓ 658ms | ✓ 931ms | ✓ 740ms | ✓ 834ms | ✓ 656ms | http |
| 193.233.22.29:10808 | ✓ 692ms | 否 | ✓ 257ms | ✓ 1781ms | ✓ 852ms | http |
| 114.237.77.241:1080 | ✓ 964ms | ✓ 1132ms | ✓ 938ms | ✓ 1311ms | ✓ 1002ms | http |
| 121.230.8.34:1080 | ✓ 982ms | ✓ 1183ms | ✓ 945ms | ✓ 1213ms | ✓ 1030ms | http |
| 45.136.130.177:8448 | ✓ 818ms | ✓ 1278ms | ✓ 1737ms | ✓ 827ms | ✓ 748ms | http |
| 183.3.221.130:3128 | ✓ 931ms | ✓ 1258ms | ✓ 1044ms | ✓ 1176ms | ✓ 959ms | http |
| 101.34.21.55:90 | ✓ 1149ms | 否 | ✓ 1028ms | ✓ 1844ms | 否 | http |
| 121.230.9.26:1080 | 否 | ✓ 1571ms | ✓ 1125ms | ✓ 1503ms | ✓ 1108ms | http |
| 121.230.8.251:1080 | ✓ 1314ms | ✓ 1715ms | ✓ 1563ms | ✓ 1618ms | ✓ 1268ms | http |
| 134.209.153.66:3128 | ✓ 1175ms | 否 | ✓ 1651ms | ✓ 1528ms | ✓ 1288ms | http |
| 114.237.77.219:1080 | 否 | 否 | ✓ 1987ms | ✓ 1251ms | ✓ 906ms | http |
| 165.22.78.128:8080 | ✓ 951ms | ✓ 1920ms | 否 | 否 | ✓ 1849ms | http |
| 103.67.46.225:3125 | ✓ 1838ms | 否 | ✓ 1706ms | 否 | ✓ 1596ms | http |
| 103.113.70.189:1081 | ✓ 446ms | 否 | ✓ 258ms | ✓ 1396ms | ✓ 1039ms | http |
| 121.230.9.252:1080 | ✓ 1243ms | ✓ 1862ms | ✓ 1510ms | ✓ 1792ms | ✓ 1139ms | http |
| 20.120.225.109:3128 | ✓ 789ms | ✓ 1182ms | ✓ 696ms | ✓ 1117ms | ✓ 812ms | http |
| 150.249.255.91:3128 | 否 | ✓ 937ms | ✓ 725ms | ✓ 927ms | ✓ 705ms | http |
| 103.82.23.118:5247 | ✓ 1487ms | ✓ 1772ms | ✓ 1147ms | ✓ 1547ms | ✓ 1702ms | http |
| 103.84.95.54:7890 | 否 | 否 | ✓ 1502ms | ✓ 1836ms | ✓ 797ms | http |

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
