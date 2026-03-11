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

最后更新：2026-03-11 16:47:27 UTC（2026-03-12 00:47:27 UTC+8）

**代理总数：77**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 77 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 77 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 45.136.130.175:8443 | ✓ 388ms | ✓ 696ms | ✓ 1252ms | ✓ 707ms | ✓ 635ms | http |
| 45.136.131.47:8443 | ✓ 388ms | ✓ 1624ms | ✓ 324ms | ✓ 706ms | ✓ 749ms | http |
| 45.136.131.63:8443 | ✓ 387ms | ✓ 1337ms | ✓ 611ms | ✓ 1125ms | ✓ 516ms | http |
| 45.136.130.188:8443 | ✓ 387ms | ✓ 1009ms | ✓ 942ms | ✓ 691ms | ✓ 1423ms | http |
| 1.231.81.166:3128 | 否 | ✓ 1265ms | ✓ 1073ms | ✓ 1159ms | ✓ 1054ms | http |
| 171.251.172.78:5110 | 否 | 否 | ✓ 1677ms | ✓ 1550ms | ✓ 1349ms | http |
| 171.251.172.78:5105 | ✓ 1884ms | 否 | ✓ 1818ms | ✓ 1503ms | ✓ 1377ms | http |
| 171.251.172.78:5108 | ✓ 1891ms | 否 | ✓ 1981ms | ✓ 1520ms | ✓ 1502ms | http |
| 101.47.73.135:3128 | ✓ 1003ms | 否 | ✓ 1699ms | 否 | ✓ 1030ms | http |
| 217.76.245.80:999 | ✓ 1064ms | 否 | ✓ 1160ms | ✓ 1548ms | ✓ 1178ms | http |
| 152.42.213.210:8080 | ✓ 826ms | 否 | ✓ 1150ms | ✓ 1128ms | ✓ 1108ms | http |
| 103.84.95.54:7890 | 否 | 否 | ✓ 656ms | ✓ 1860ms | ✓ 1582ms | http |
| 185.191.236.162:3128 | ✓ 835ms | ✓ 1744ms | 否 | 否 | ✓ 1718ms | http |
| 95.3.9.78:3128 | ✓ 876ms | 否 | 否 | ✓ 1809ms | ✓ 1386ms | http |
| 115.231.181.40:8128 | ✓ 1208ms | 否 | ✓ 1509ms | ✓ 1567ms | 否 | http |
| 190.9.109.198:999 | 否 | 否 | ✓ 1102ms | ✓ 1391ms | ✓ 1376ms | http |
| 59.46.216.131:30001 | ✓ 1029ms | 否 | ✓ 1715ms | ✓ 1820ms | ✓ 1118ms | http |
| 111.48.191.1:7890 | ✓ 727ms | ✓ 920ms | ✓ 843ms | ✓ 943ms | ✓ 720ms | http |
| 165.227.5.10:8888 | ✓ 647ms | 否 | ✓ 737ms | ✓ 1925ms | ✓ 1880ms | http |
| 45.136.130.191:8443 | ✓ 428ms | ✓ 1074ms | ✓ 355ms | ✓ 951ms | ✓ 518ms | http |
| 120.92.212.16:8890 | ✓ 1236ms | ✓ 1404ms | ✓ 1310ms | 否 | 否 | http |
| 81.70.169.194:80 | ✓ 1856ms | ✓ 1433ms | ✓ 1339ms | ✓ 1303ms | 否 | http |
| 39.104.201.40:7890 | 否 | ✓ 1491ms | ✓ 979ms | 否 | ✓ 1411ms | http |
| 116.80.96.106:3172 | ✓ 1687ms | 否 | 否 | ✓ 1887ms | ✓ 1664ms | http |
| 138.124.53.25:7443 | ✓ 603ms | 否 | ✓ 1380ms | ✓ 1585ms | 否 | http |
| 211.171.114.154:3128 | ✓ 1290ms | ✓ 1641ms | 否 | ✓ 1666ms | ✓ 1587ms | http |
| 91.107.141.42:8081 | 否 | 否 | ✓ 1891ms | ✓ 1853ms | ✓ 1792ms | http |
| 107.173.0.178:1080 | ✓ 482ms | 否 | ✓ 913ms | ✓ 1567ms | ✓ 1318ms | http |
| 46.183.25.8:443 | ✓ 1074ms | 否 | ✓ 146ms | ✓ 744ms | 否 | http |
| 152.42.213.210:443 | ✓ 799ms | 否 | ✓ 988ms | ✓ 1175ms | ✓ 837ms | http |
| 47.101.149.27:9010 | 否 | ✓ 1421ms | ✓ 1321ms | ✓ 1671ms | ✓ 1242ms | http |
| 120.92.212.16:7890 | ✓ 1492ms | 否 | 否 | ✓ 1367ms | ✓ 1069ms | http |
| 190.212.131.238:3128 | ✓ 1461ms | 否 | ✓ 1563ms | 否 | ✓ 1705ms | http |
| 101.43.255.96:80 | ✓ 1309ms | 否 | 否 | ✓ 1285ms | ✓ 1068ms | http |
| 210.223.44.230:3128 | 否 | 否 | ✓ 1223ms | ✓ 1684ms | ✓ 1845ms | http |
| 114.55.226.123:10086 | ✓ 1031ms | 否 | ✓ 998ms | ✓ 1253ms | ✓ 1077ms | http |
| 205.209.118.30:3138 | ✓ 1899ms | 否 | 否 | ✓ 1293ms | ✓ 1007ms | http |
| 107.172.125.217:3128 | ✓ 103ms | 否 | ✓ 725ms | ✓ 796ms | ✓ 681ms | http |
| 34.101.184.164:3128 | ✓ 1696ms | 否 | ✓ 1293ms | 否 | ✓ 968ms | http |
| 154.9.235.189:7890 | ✓ 687ms | 否 | ✓ 428ms | ✓ 1748ms | ✓ 1969ms | http |
| 194.213.18.200:443 | ✓ 667ms | 否 | 否 | ✓ 1144ms | ✓ 1112ms | http |
| 202.155.12.161:443 | ✓ 1664ms | 否 | 否 | ✓ 1091ms | ✓ 902ms | http |
| 95.3.9.78:8080 | ✓ 1091ms | 否 | ✓ 804ms | ✓ 1822ms | 否 | http |
| 171.251.172.78:5107 | 否 | 否 | ✓ 1976ms | ✓ 1474ms | ✓ 1398ms | http |
| 35.225.22.61:80 | ✓ 581ms | 否 | ✓ 443ms | 否 | ✓ 904ms | http |
| 159.223.42.219:3128 | ✓ 745ms | 否 | 否 | ✓ 1269ms | ✓ 836ms | http |
| 171.251.172.78:5104 | ✓ 1704ms | 否 | ✓ 1261ms | ✓ 1582ms | ✓ 1333ms | http |
| 45.136.198.40:3128 | ✓ 821ms | 否 | ✓ 830ms | 否 | ✓ 1783ms | http |
| 85.208.108.43:2094 | ✓ 1290ms | 否 | 否 | ✓ 1786ms | ✓ 1876ms | http |
| 43.167.227.161:1080 | ✓ 1364ms | 否 | ✓ 797ms | ✓ 782ms | ✓ 591ms | http |
| 167.172.67.118:8080 | ✓ 747ms | 否 | ✓ 1331ms | 否 | ✓ 830ms | http |
| 171.251.172.78:5102 | ✓ 1949ms | 否 | ✓ 1432ms | ✓ 1478ms | ✓ 1325ms | http |
| 171.251.172.78:5106 | ✓ 1721ms | 否 | ✓ 1682ms | ✓ 1565ms | ✓ 1332ms | http |
| 116.80.49.167:3172 | ✓ 1648ms | 否 | ✓ 1600ms | 否 | ✓ 1853ms | http |
| 101.32.244.83:8080 | ✓ 989ms | 否 | ✓ 935ms | ✓ 1374ms | ✓ 1212ms | http |
| 121.43.196.210:8222 | ✓ 942ms | ✓ 1012ms | ✓ 971ms | ✓ 1163ms | ✓ 908ms | http |
| 121.43.196.213:8222 | ✓ 1038ms | ✓ 1098ms | ✓ 956ms | ✓ 1146ms | ✓ 991ms | http |
| 178.236.245.17:3128 | ✓ 914ms | ✓ 1944ms | ✓ 1356ms | 否 | ✓ 1747ms | http |
| 150.249.255.91:3128 | ✓ 1045ms | ✓ 1074ms | ✓ 1786ms | ✓ 1293ms | ✓ 963ms | http |
| 171.251.172.78:5109 | 否 | 否 | ✓ 1500ms | ✓ 1497ms | ✓ 1348ms | http |
| 14.225.222.164:7890 | ✓ 1957ms | 否 | ✓ 1420ms | 否 | ✓ 1590ms | http |
| 106.117.208.101:7890 | ✓ 1932ms | ✓ 1316ms | ✓ 1706ms | ✓ 1285ms | 否 | http |
| 107.173.52.58:7890 | ✓ 439ms | 否 | 否 | ✓ 1149ms | ✓ 904ms | http |
| 113.160.132.26:8080 | ✓ 1798ms | ✓ 1339ms | ✓ 1057ms | ✓ 1222ms | ✓ 1007ms | http |
| 47.77.193.180:1080 | ✓ 201ms | ✓ 887ms | ✓ 256ms | ✓ 797ms | ✓ 666ms | http |
| 45.136.130.239:8443 | 否 | ✓ 1407ms | 否 | ✓ 787ms | ✓ 520ms | http |
| 172.212.68.37:3128 | ✓ 1766ms | 否 | ✓ 243ms | ✓ 1729ms | ✓ 1108ms | http |
| 103.39.51.190:8080 | 否 | 否 | ✓ 1832ms | ✓ 1458ms | ✓ 1387ms | http |
| 168.235.110.63:3128 | ✓ 583ms | ✓ 1436ms | ✓ 357ms | 否 | ✓ 1319ms | http |
| 1.225.116.115:1080 | ✓ 1893ms | 否 | ✓ 1482ms | 否 | ✓ 1104ms | http |
| 116.80.96.105:3172 | ✓ 1505ms | 否 | 否 | ✓ 1801ms | ✓ 1613ms | http |
| 8.219.97.248:80 | 否 | 否 | ✓ 1031ms | ✓ 1750ms | ✓ 1269ms | http |
| 62.113.119.14:8080 | 否 | 否 | ✓ 1199ms | ✓ 1632ms | ✓ 1209ms | http |
| 152.70.98.46:8888 | ✓ 1200ms | ✓ 1370ms | 否 | 否 | ✓ 1342ms | http |
| 61.52.131.172:8443 | ✓ 849ms | ✓ 1172ms | ✓ 1005ms | ✓ 1251ms | ✓ 924ms | http |
| 89.251.9.11:3128 | ✓ 671ms | ✓ 1426ms | ✓ 1755ms | ✓ 1166ms | 否 | http |
| 103.183.10.172:3125 | 否 | 否 | ✓ 1673ms | ✓ 1652ms | ✓ 1425ms | http |

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
