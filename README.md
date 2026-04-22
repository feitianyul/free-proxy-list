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

最后更新：2026-04-22 21:54:23 UTC（2026-04-23 05:54:23 UTC+8）

**代理总数：82**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 82 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 82 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 1.231.81.166:3128 | ✓ 1508ms | ✓ 971ms | ✓ 951ms | ✓ 840ms | ✓ 671ms | http |
| 218.108.131.186:17890 | ✓ 833ms | ✓ 980ms | ✓ 828ms | ✓ 1052ms | ✓ 852ms | http |
| 113.160.132.26:8080 | ✓ 1494ms | ✓ 1496ms | 否 | ✓ 1212ms | ✓ 941ms | http |
| 152.42.208.139:8118 | ✓ 698ms | 否 | ✓ 1557ms | ✓ 1030ms | ✓ 1881ms | http |
| 223.84.151.86:30005 | 否 | 否 | ✓ 1274ms | ✓ 1838ms | ✓ 1767ms | http |
| 78.11.96.22:8888 | ✓ 1264ms | ✓ 1770ms | ✓ 1954ms | ✓ 1740ms | 否 | http |
| 115.231.181.40:8128 | ✓ 877ms | ✓ 995ms | ✓ 850ms | ✓ 1223ms | 否 | http |
| 120.92.108.86:7890 | ✓ 1254ms | 否 | ✓ 1146ms | ✓ 1793ms | ✓ 1667ms | http |
| 177.93.132.244:3128 | ✓ 1759ms | 否 | ✓ 980ms | 否 | ✓ 1888ms | http |
| 91.99.15.45:2095 | ✓ 1139ms | ✓ 1783ms | ✓ 1731ms | 否 | 否 | http |
| 38.180.192.119:3128 | ✓ 1431ms | ✓ 1488ms | ✓ 748ms | ✓ 1182ms | ✓ 830ms | http |
| 210.223.44.230:3128 | ✓ 1776ms | ✓ 927ms | ✓ 1074ms | ✓ 1887ms | ✓ 1657ms | http |
| 84.47.150.125:1080 | ✓ 1111ms | 否 | ✓ 1680ms | 否 | ✓ 1838ms | http |
| 46.101.95.183:8888 | ✓ 1994ms | 否 | 否 | ✓ 1987ms | ✓ 1552ms | http |
| 85.190.99.143:443 | ✓ 1678ms | 否 | ✓ 1727ms | 否 | ✓ 1922ms | http |
| 35.225.22.61:80 | ✓ 665ms | ✓ 1384ms | ✓ 1027ms | ✓ 1255ms | 否 | http |
| 172.233.49.176:9658 | ✓ 1516ms | 否 | ✓ 858ms | 否 | ✓ 1680ms | http |
| 45.140.147.155:1081 | ✓ 1972ms | 否 | ✓ 874ms | 否 | ✓ 1336ms | http |
| 103.39.51.190:8090 | ✓ 1269ms | 否 | 否 | ✓ 1521ms | ✓ 1867ms | http |
| 122.224.198.218:808 | ✓ 1886ms | 否 | ✓ 1910ms | 否 | ✓ 1974ms | http |
| 208.87.243.199:7878 | ✓ 1068ms | 否 | ✓ 477ms | 否 | ✓ 1909ms | http |
| 59.46.216.131:30001 | 否 | ✓ 1318ms | 否 | ✓ 1333ms | ✓ 1146ms | http |
| 34.96.238.40:8080 | ✓ 978ms | 否 | ✓ 791ms | ✓ 1093ms | 否 | http |
| 62.113.119.14:8080 | ✓ 1471ms | ✓ 1718ms | ✓ 1112ms | ✓ 1617ms | ✓ 1256ms | http |
| 45.129.141.143:3128 | ✓ 1101ms | 否 | ✓ 1983ms | 否 | ✓ 1893ms | http |
| 130.61.174.200:1080 | ✓ 1601ms | 否 | ✓ 1601ms | ✓ 1741ms | 否 | http |
| 45.140.147.82:1082 | ✓ 1280ms | 否 | ✓ 1907ms | 否 | ✓ 1793ms | http |
| 212.58.132.5:8888 | ✓ 1780ms | 否 | ✓ 1324ms | ✓ 1890ms | ✓ 1277ms | http |
| 42.101.8.101:8888 | ✓ 1718ms | ✓ 1581ms | ✓ 1273ms | 否 | ✓ 1028ms | http |
| 217.76.245.80:999 | ✓ 1332ms | ✓ 1372ms | ✓ 1469ms | ✓ 1776ms | ✓ 1379ms | http |
| 121.43.196.210:8222 | ✓ 857ms | ✓ 1018ms | ✓ 871ms | ✓ 1061ms | ✓ 879ms | http |
| 170.9.253.20:8888 | ✓ 957ms | ✓ 1740ms | 否 | ✓ 1837ms | ✓ 1146ms | http |
| 89.208.106.138:10808 | ✓ 1295ms | ✓ 1818ms | 否 | ✓ 1767ms | 否 | http |
| 47.85.51.197:1080 | ✓ 791ms | ✓ 1565ms | ✓ 858ms | ✓ 1184ms | ✓ 1433ms | http |
| 193.122.96.242:3128 | ✓ 1393ms | 否 | ✓ 1693ms | ✓ 1549ms | ✓ 998ms | http |
| 45.153.231.229:8080 | ✓ 1428ms | 否 | ✓ 1075ms | 否 | ✓ 1868ms | http |
| 194.150.220.163:1082 | 否 | 否 | ✓ 1652ms | ✓ 1831ms | ✓ 1595ms | http |
| 34.71.229.255:3128 | ✓ 1619ms | ✓ 1717ms | ✓ 1870ms | ✓ 1552ms | ✓ 1276ms | http |
| 45.140.147.82:1081 | ✓ 748ms | ✓ 1476ms | ✓ 1201ms | ✓ 1894ms | ✓ 1525ms | http |
| 8.209.238.110:47701 | ✓ 507ms | ✓ 906ms | ✓ 602ms | ✓ 808ms | ✓ 635ms | http |
| 168.138.202.218:3128 | ✓ 1454ms | ✓ 1187ms | ✓ 1391ms | ✓ 1064ms | ✓ 718ms | http |
| 121.230.8.237:1080 | ✓ 1134ms | ✓ 1486ms | ✓ 924ms | ✓ 1228ms | ✓ 1163ms | http |
| 103.187.146.151:3128 | ✓ 771ms | 否 | ✓ 998ms | ✓ 1076ms | ✓ 940ms | http |
| 120.92.212.16:7890 | ✓ 894ms | ✓ 1907ms | ✓ 1013ms | ✓ 1360ms | ✓ 1402ms | http |
| 182.204.176.114:1080 | ✓ 1025ms | ✓ 1446ms | ✓ 1214ms | ✓ 1787ms | ✓ 1124ms | http |
| 120.92.212.16:8890 | ✓ 863ms | 否 | 否 | ✓ 1325ms | ✓ 1224ms | http |
| 121.230.8.49:1080 | ✓ 1024ms | ✓ 1262ms | ✓ 1100ms | ✓ 1484ms | ✓ 993ms | http |
| 103.240.6.22:16498 | ✓ 1306ms | 否 | ✓ 1305ms | ✓ 1377ms | ✓ 1644ms | http |
| 101.32.244.83:8080 | ✓ 1509ms | ✓ 1392ms | ✓ 894ms | ✓ 1367ms | ✓ 1192ms | http |
| 121.43.196.213:8222 | ✓ 926ms | ✓ 1008ms | ✓ 810ms | ✓ 1065ms | ✓ 863ms | http |
| 85.214.39.41:3128 | ✓ 1114ms | 否 | ✓ 1770ms | 否 | ✓ 1891ms | http |
| 14.143.222.113:57788 | ✓ 1169ms | 否 | ✓ 1726ms | ✓ 1890ms | 否 | http |
| 77.110.113.24:40000 | ✓ 1119ms | ✓ 1974ms | 否 | 否 | ✓ 1843ms | http |
| 152.42.177.32:8888 | ✓ 897ms | 否 | ✓ 1816ms | ✓ 1219ms | ✓ 1227ms | http |
| 45.59.122.132:80 | ✓ 1608ms | 否 | ✓ 1208ms | 否 | ✓ 1480ms | http |
| 150.107.140.238:3128 | ✓ 822ms | 否 | ✓ 881ms | 否 | ✓ 979ms | http |
| 45.186.6.104:3128 | ✓ 1728ms | ✓ 1647ms | ✓ 1858ms | 否 | 否 | http |
| 8.219.97.248:80 | ✓ 1716ms | 否 | 否 | ✓ 1722ms | ✓ 1295ms | http |
| 116.171.106.26:3443 | ✓ 1543ms | ✓ 1411ms | ✓ 1382ms | ✓ 1737ms | ✓ 1445ms | http |
| 202.141.161.53:10808 | 否 | ✓ 1771ms | ✓ 1124ms | ✓ 1167ms | 否 | http |
| 150.249.255.91:3128 | 否 | ✓ 1204ms | ✓ 497ms | 否 | ✓ 634ms | http |
| 20.205.16.149:3128 | ✓ 617ms | ✓ 1295ms | ✓ 916ms | ✓ 881ms | ✓ 1828ms | http |
| 8.219.195.129:1080 | ✓ 681ms | ✓ 1635ms | ✓ 928ms | ✓ 1010ms | 否 | http |
| 152.32.132.190:7890 | ✓ 633ms | ✓ 1819ms | ✓ 774ms | ✓ 1288ms | ✓ 1414ms | http |
| 15.204.233.75:3128 | ✓ 1269ms | ✓ 1435ms | ✓ 1776ms | ✓ 1834ms | ✓ 1429ms | http |
| 47.110.42.192:9003 | ✓ 1411ms | ✓ 1244ms | ✓ 1283ms | ✓ 1701ms | ✓ 1269ms | http |
| 84.247.171.137:3128 | ✓ 1713ms | ✓ 1657ms | 否 | 否 | ✓ 1726ms | http |
| 36.92.104.123:8000 | ✓ 1787ms | 否 | ✓ 1625ms | ✓ 1462ms | ✓ 1461ms | http |
| 103.160.68.201:8085 | 否 | 否 | ✓ 1211ms | ✓ 1399ms | ✓ 1451ms | http |
| 104.248.243.244:3128 | ✓ 1147ms | 否 | ✓ 1609ms | ✓ 1747ms | ✓ 1237ms | http |
| 45.76.207.177:40000 | ✓ 763ms | 否 | ✓ 906ms | ✓ 1015ms | ✓ 906ms | http |
| 138.118.104.33:8080 | ✓ 1065ms | 否 | ✓ 1770ms | ✓ 1742ms | ✓ 1423ms | http |
| 108.181.201.118:1234 | ✓ 924ms | 否 | ✓ 1075ms | ✓ 1156ms | 否 | http |
| 195.26.224.49:3128 | 否 | 否 | ✓ 796ms | ✓ 1743ms | ✓ 1239ms | http |
| 147.45.186.28:3128 | ✓ 1073ms | ✓ 1984ms | ✓ 1038ms | ✓ 1904ms | 否 | http |
| 47.84.73.61:1080 | ✓ 815ms | ✓ 1599ms | ✓ 817ms | ✓ 988ms | ✓ 804ms | http |
| 114.55.226.123:10086 | ✓ 1143ms | 否 | ✓ 1067ms | ✓ 1280ms | 否 | http |
| 61.52.131.172:8443 | ✓ 855ms | ✓ 1102ms | ✓ 911ms | ✓ 1128ms | ✓ 917ms | http |
| 121.230.8.144:1080 | ✓ 1023ms | ✓ 1699ms | ✓ 973ms | 否 | 否 | http |
| 147.45.60.34:1082 | 否 | 否 | ✓ 287ms | ✓ 1491ms | ✓ 986ms | http |
| 103.39.51.207:8080 | ✓ 1184ms | 否 | 否 | ✓ 1321ms | ✓ 1260ms | http |
| 213.220.3.234:20573 | ✓ 1253ms | ✓ 1723ms | ✓ 1790ms | 否 | 否 | http |

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
