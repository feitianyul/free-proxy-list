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

最后更新：2026-04-17 08:47:08 UTC（2026-04-17 16:47:08 UTC+8）

**代理总数：73**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 73 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 73 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 34.96.238.40:8080 | 否 | ✓ 1203ms | ✓ 914ms | 否 | ✓ 972ms | http |
| 149.51.42.10:3128 | ✓ 1445ms | ✓ 1778ms | 否 | ✓ 1478ms | 否 | http |
| 185.114.73.2:1080 | ✓ 1203ms | 否 | ✓ 1699ms | ✓ 1635ms | ✓ 1424ms | http |
| 1.231.81.166:3128 | 否 | ✓ 1952ms | ✓ 1457ms | ✓ 1306ms | ✓ 993ms | http |
| 113.160.132.26:8080 | ✓ 1509ms | ✓ 1415ms | ✓ 1764ms | ✓ 1390ms | ✓ 1227ms | http |
| 101.32.243.189:80 | 否 | 否 | ✓ 1224ms | ✓ 1472ms | ✓ 1583ms | http |
| 218.108.131.186:17890 | ✓ 1826ms | ✓ 1527ms | ✓ 862ms | ✓ 1523ms | 否 | http |
| 20.127.128.70:8080 | ✓ 1260ms | 否 | ✓ 913ms | 否 | ✓ 1861ms | http |
| 157.230.178.216:8088 | ✓ 323ms | ✓ 1386ms | ✓ 498ms | ✓ 1193ms | ✓ 1140ms | http |
| 94.131.118.129:1081 | ✓ 618ms | ✓ 1485ms | ✓ 1782ms | 否 | 否 | http |
| 147.45.214.210:1080 | ✓ 907ms | ✓ 1686ms | 否 | 否 | ✓ 1387ms | http |
| 177.93.132.244:3128 | ✓ 1191ms | 否 | ✓ 1882ms | 否 | ✓ 1647ms | http |
| 82.114.228.67:1080 | ✓ 848ms | 否 | 否 | ✓ 1919ms | ✓ 1281ms | http |
| 168.144.75.9:3128 | ✓ 1588ms | 否 | ✓ 1583ms | ✓ 1937ms | ✓ 1558ms | http |
| 149.51.42.10:8080 | ✓ 803ms | ✓ 1462ms | 否 | ✓ 1921ms | 否 | http |
| 138.124.99.216:8888 | ✓ 770ms | 否 | 否 | ✓ 1709ms | ✓ 1761ms | http |
| 188.246.224.49:7890 | ✓ 1255ms | 否 | ✓ 1261ms | 否 | ✓ 1200ms | http |
| 42.101.8.101:8888 | ✓ 1163ms | ✓ 1383ms | ✓ 1073ms | ✓ 1366ms | ✓ 1091ms | http |
| 16.62.127.160:3128 | ✓ 1422ms | 否 | ✓ 975ms | 否 | ✓ 1753ms | http |
| 185.138.116.150:8080 | ✓ 731ms | 否 | ✓ 1603ms | 否 | ✓ 1734ms | http |
| 15.161.131.175:3129 | ✓ 1314ms | 否 | ✓ 1700ms | ✓ 1890ms | 否 | http |
| 51.95.13.205:12 | ✓ 1318ms | 否 | ✓ 1761ms | 否 | ✓ 1880ms | http |
| 202.141.161.53:10808 | 否 | 否 | ✓ 1114ms | ✓ 1896ms | ✓ 1104ms | http |
| 149.104.4.88:10809 | ✓ 889ms | 否 | ✓ 1280ms | ✓ 827ms | ✓ 676ms | http |
| 212.58.132.5:8888 | ✓ 1202ms | 否 | 否 | ✓ 1615ms | ✓ 1307ms | http |
| 43.132.188.134:443 | ✓ 641ms | 否 | 否 | ✓ 1957ms | ✓ 1718ms | http |
| 162.240.154.26:3128 | 否 | 否 | ✓ 1759ms | ✓ 1991ms | ✓ 1858ms | http |
| 116.80.47.79:3128 | ✓ 1499ms | 否 | 否 | ✓ 1814ms | ✓ 1649ms | http |
| 8.219.195.129:1080 | ✓ 1489ms | 否 | ✓ 899ms | ✓ 1066ms | ✓ 833ms | http |
| 34.71.229.255:3128 | 否 | 否 | ✓ 1181ms | ✓ 1838ms | ✓ 1384ms | http |
| 34.246.223.187:3128 | ✓ 1996ms | 否 | ✓ 1176ms | 否 | ✓ 1456ms | http |
| 117.236.124.166:3128 | ✓ 1897ms | 否 | ✓ 1711ms | 否 | ✓ 1624ms | http |
| 103.113.70.189:1081 | ✓ 780ms | 否 | ✓ 1283ms | ✓ 1421ms | ✓ 882ms | http |
| 94.131.118.39:1081 | ✓ 772ms | 否 | ✓ 1751ms | 否 | ✓ 1913ms | http |
| 84.47.150.125:1080 | ✓ 863ms | 否 | ✓ 1950ms | 否 | ✓ 1758ms | http |
| 34.101.184.164:3128 | ✓ 1620ms | 否 | ✓ 1738ms | ✓ 1569ms | ✓ 1205ms | http |
| 94.158.219.111:3128 | ✓ 1350ms | 否 | ✓ 1593ms | 否 | ✓ 1917ms | http |
| 103.138.70.165:3129 | ✓ 1824ms | 否 | 否 | ✓ 1895ms | ✓ 1423ms | http |
| 59.46.216.131:30001 | 否 | ✓ 1398ms | ✓ 1323ms | 否 | ✓ 1303ms | http |
| 209.126.84.232:8888 | 否 | 否 | ✓ 1667ms | ✓ 1587ms | ✓ 1794ms | http |
| 113.176.92.71:3128 | 否 | 否 | ✓ 1552ms | ✓ 1239ms | ✓ 987ms | http |
| 27.71.24.102:3128 | ✓ 1999ms | 否 | ✓ 1245ms | ✓ 1066ms | ✓ 1024ms | http |
| 103.87.202.198:1111 | ✓ 1348ms | 否 | ✓ 1638ms | 否 | ✓ 1409ms | http |
| 62.113.119.14:8080 | ✓ 778ms | ✓ 1705ms | ✓ 824ms | 否 | 否 | http |
| 52.16.215.4:7593 | ✓ 1228ms | 否 | ✓ 720ms | ✓ 1860ms | ✓ 1549ms | http |
| 147.45.186.28:3128 | 否 | 否 | ✓ 946ms | ✓ 1688ms | ✓ 1932ms | http |
| 116.80.96.121:3128 | 否 | 否 | ✓ 1528ms | ✓ 1832ms | ✓ 1661ms | http |
| 139.159.97.82:10900 | 否 | 否 | ✓ 1712ms | ✓ 1458ms | ✓ 1359ms | http |
| 113.192.12.37:8181 | ✓ 1778ms | 否 | ✓ 1229ms | ✓ 1582ms | ✓ 1414ms | http |
| 103.93.93.95:8181 | ✓ 1785ms | 否 | ✓ 1648ms | ✓ 1510ms | ✓ 1479ms | http |
| 157.10.97.12:8080 | ✓ 1800ms | 否 | ✓ 1482ms | ✓ 1784ms | ✓ 1475ms | http |
| 126.209.18.142:8082 | ✓ 1797ms | 否 | 否 | ✓ 1533ms | ✓ 1536ms | http |
| 35.225.22.61:80 | ✓ 909ms | 否 | ✓ 583ms | ✓ 1586ms | ✓ 1069ms | http |
| 103.85.113.66:9999 | ✓ 819ms | ✓ 1678ms | ✓ 1257ms | 否 | 否 | http |
| 218.153.163.186:8475 | 否 | ✓ 1919ms | ✓ 1965ms | ✓ 1809ms | ✓ 1131ms | http |
| 196.206.98.64:1221 | ✓ 1746ms | 否 | ✓ 1415ms | 否 | ✓ 1851ms | http |
| 52.59.51.29:30149 | ✓ 1540ms | 否 | ✓ 1720ms | 否 | ✓ 1871ms | http |
| 223.84.151.86:30005 | ✓ 1530ms | ✓ 1512ms | ✓ 1113ms | ✓ 1425ms | 否 | http |
| 117.122.240.82:3338 | ✓ 1058ms | ✓ 1795ms | 否 | ✓ 1668ms | ✓ 1324ms | http |
| 220.197.44.36:3128 | ✓ 1264ms | ✓ 1396ms | ✓ 1136ms | ✓ 1442ms | ✓ 1235ms | http |
| 13.38.217.179:3128 | ✓ 1201ms | 否 | ✓ 732ms | 否 | ✓ 1709ms | http |
| 45.12.151.226:2829 | ✓ 1163ms | 否 | ✓ 1412ms | 否 | ✓ 1876ms | http |
| 103.113.70.189:1082 | ✓ 1917ms | ✓ 1656ms | ✓ 420ms | 否 | ✓ 1176ms | http |
| 210.223.44.230:3128 | 否 | 否 | ✓ 1019ms | ✓ 1595ms | ✓ 1401ms | http |
| 45.140.147.82:1081 | ✓ 1303ms | ✓ 1705ms | ✓ 862ms | ✓ 1581ms | ✓ 1316ms | http |
| 103.18.77.14:1111 | ✓ 1941ms | 否 | 否 | ✓ 1315ms | ✓ 1651ms | http |
| 61.52.131.172:8443 | ✓ 1708ms | ✓ 1167ms | ✓ 950ms | ✓ 1110ms | ✓ 967ms | http |
| 168.138.202.218:3128 | ✓ 1776ms | 否 | ✓ 1104ms | ✓ 1782ms | ✓ 776ms | http |
| 51.16.244.65:3129 | ✓ 1313ms | 否 | ✓ 1356ms | 否 | ✓ 1806ms | http |
| 36.141.21.200:7890 | ✓ 1000ms | ✓ 1244ms | ✓ 1025ms | 否 | ✓ 979ms | http |
| 130.61.30.221:8080 | ✓ 1201ms | 否 | ✓ 1860ms | ✓ 1568ms | 否 | http |
| 116.58.161.203:26021 | ✓ 1752ms | 否 | ✓ 1770ms | ✓ 1146ms | ✓ 1216ms | http |
| 160.250.5.22:1 | 否 | 否 | ✓ 1509ms | ✓ 1199ms | ✓ 1007ms | http |

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
