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

最后更新：2026-05-01 12:01:17 UTC（2026-05-01 20:01:17 UTC+8）

**代理总数：66**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 66 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 66 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 1.231.81.166:3128 | ✓ 912ms | ✓ 1182ms | ✓ 1656ms | ✓ 1221ms | ✓ 1265ms | http |
| 206.206.126.177:2412 | ✓ 1445ms | 否 | ✓ 864ms | ✓ 1138ms | ✓ 897ms | http |
| 113.160.132.26:8080 | ✓ 1598ms | 否 | 否 | ✓ 1475ms | ✓ 1157ms | http |
| 45.167.124.71:999 | ✓ 858ms | 否 | ✓ 1270ms | ✓ 1637ms | ✓ 1306ms | http |
| 46.101.95.183:8888 | ✓ 716ms | 否 | ✓ 914ms | 否 | ✓ 1170ms | http |
| 34.101.184.164:3128 | ✓ 1784ms | 否 | ✓ 1203ms | ✓ 1463ms | ✓ 1144ms | http |
| 8.219.97.248:80 | ✓ 1074ms | 否 | ✓ 1267ms | 否 | ✓ 1644ms | http |
| 194.150.220.163:1082 | ✓ 760ms | ✓ 1827ms | ✓ 972ms | ✓ 1670ms | ✓ 1297ms | http |
| 103.157.200.126:3128 | ✓ 1594ms | 否 | ✓ 1898ms | ✓ 1708ms | ✓ 1900ms | http |
| 103.70.114.149:3128 | ✓ 1643ms | 否 | ✓ 1613ms | ✓ 1869ms | ✓ 1749ms | http |
| 43.167.237.94:3128 | ✓ 1402ms | 否 | ✓ 1467ms | ✓ 953ms | ✓ 728ms | http |
| 47.85.51.197:1080 | ✓ 1189ms | 否 | ✓ 1144ms | 否 | ✓ 735ms | http |
| 115.231.181.40:8128 | ✓ 1590ms | ✓ 1186ms | ✓ 1024ms | ✓ 1357ms | 否 | http |
| 120.92.108.86:7890 | ✓ 1304ms | 否 | ✓ 1440ms | ✓ 1939ms | 否 | http |
| 80.92.204.47:1081 | ✓ 503ms | ✓ 1877ms | ✓ 1073ms | ✓ 1689ms | ✓ 1270ms | http |
| 62.113.119.14:8080 | ✓ 717ms | ✓ 1540ms | ✓ 1276ms | 否 | 否 | http |
| 130.61.174.200:1080 | ✓ 752ms | ✓ 1250ms | ✓ 1964ms | 否 | 否 | http |
| 168.110.52.228:3128 | 否 | ✓ 1968ms | 否 | ✓ 1352ms | ✓ 1230ms | http |
| 34.96.238.40:8080 | ✓ 1223ms | 否 | 否 | ✓ 1780ms | ✓ 1160ms | http |
| 218.108.131.186:17890 | ✓ 937ms | ✓ 1615ms | ✓ 964ms | 否 | ✓ 1002ms | http |
| 212.58.132.5:8888 | ✓ 1573ms | 否 | ✓ 1453ms | ✓ 1596ms | ✓ 1295ms | http |
| 120.92.212.16:7890 | ✓ 1085ms | ✓ 1419ms | ✓ 1610ms | 否 | 否 | http |
| 103.35.190.182:1081 | ✓ 1712ms | ✓ 969ms | ✓ 855ms | ✓ 1193ms | ✓ 947ms | http |
| 103.35.190.182:1082 | ✓ 1715ms | ✓ 971ms | ✓ 851ms | ✓ 1163ms | 否 | http |
| 152.70.91.193:40000 | ✓ 1659ms | 否 | 否 | ✓ 1414ms | ✓ 1496ms | http |
| 43.133.44.89:8888 | ✓ 999ms | ✓ 1983ms | 否 | ✓ 1135ms | ✓ 934ms | http |
| 49.7.179.70:3333 | ✓ 1295ms | ✓ 1347ms | ✓ 1733ms | 否 | 否 | http |
| 86.104.72.220:1081 | ✓ 1607ms | ✓ 1650ms | ✓ 1583ms | 否 | 否 | http |
| 120.92.212.16:8890 | ✓ 1880ms | ✓ 1657ms | ✓ 1405ms | ✓ 1634ms | ✓ 1307ms | http |
| 150.107.140.238:3128 | ✓ 1689ms | 否 | ✓ 1061ms | ✓ 1472ms | ✓ 1110ms | http |
| 210.223.44.230:3128 | ✓ 1635ms | ✓ 1547ms | 否 | 否 | ✓ 1039ms | http |
| 223.84.151.86:30005 | 否 | ✓ 1345ms | ✓ 1659ms | ✓ 1621ms | ✓ 1665ms | http |
| 86.104.72.219:1081 | ✓ 1957ms | 否 | ✓ 727ms | ✓ 1562ms | 否 | http |
| 217.77.102.18:3128 | ✓ 1116ms | ✓ 1956ms | ✓ 1366ms | 否 | ✓ 1754ms | http |
| 121.232.73.214:1080 | 否 | 否 | ✓ 1497ms | ✓ 1831ms | ✓ 1662ms | http |
| 159.223.225.118:8888 | ✓ 1084ms | 否 | ✓ 697ms | 否 | ✓ 1503ms | http |
| 154.64.232.35:8080 | 否 | ✓ 833ms | ✓ 1884ms | ✓ 962ms | ✓ 1209ms | http |
| 183.232.248.73:7890 | ✓ 1280ms | ✓ 1314ms | ✓ 1095ms | ✓ 1195ms | ✓ 974ms | http |
| 218.72.124.35:40000 | ✓ 1020ms | ✓ 1242ms | ✓ 946ms | ✓ 1568ms | ✓ 1044ms | http |
| 183.238.3.150:7897 | ✓ 1324ms | ✓ 1326ms | ✓ 1213ms | ✓ 1257ms | ✓ 999ms | http |
| 3.101.133.120:80 | ✓ 406ms | ✓ 1409ms | ✓ 462ms | ✓ 1331ms | ✓ 1182ms | http |
| 92.119.127.212:6005 | ✓ 1537ms | 否 | ✓ 1745ms | ✓ 1687ms | ✓ 1882ms | http |
| 89.43.134.23:8080 | ✓ 1588ms | ✓ 1891ms | ✓ 1993ms | 否 | 否 | http |
| 185.230.191.240:3128 | ✓ 1073ms | 否 | ✓ 1160ms | 否 | ✓ 1635ms | http |
| 15.204.151.142:3128 | ✓ 541ms | 否 | ✓ 1721ms | ✓ 1539ms | 否 | http |
| 64.188.67.154:1080 | 否 | 否 | ✓ 1201ms | ✓ 1590ms | ✓ 1252ms | http |
| 220.197.44.36:3128 | 否 | ✓ 1483ms | 否 | ✓ 1892ms | ✓ 1654ms | http |
| 59.46.216.131:30001 | 否 | ✓ 1568ms | 否 | ✓ 1458ms | ✓ 1233ms | http |
| 173.212.245.136:8888 | ✓ 1445ms | 否 | ✓ 1806ms | 否 | ✓ 1773ms | http |
| 86.104.74.110:1082 | ✓ 670ms | 否 | ✓ 1016ms | 否 | ✓ 1926ms | http |
| 86.104.74.110:1081 | ✓ 662ms | ✓ 1516ms | ✓ 1492ms | 否 | ✓ 1919ms | http |
| 45.186.6.104:3128 | ✓ 1199ms | ✓ 1902ms | ✓ 1977ms | 否 | 否 | http |
| 72.11.151.159:6005 | ✓ 859ms | 否 | ✓ 934ms | ✓ 1157ms | ✓ 803ms | http |
| 103.35.190.69:1082 | 否 | ✓ 1513ms | ✓ 225ms | ✓ 1194ms | 否 | http |
| 38.188.247.12:999 | ✓ 872ms | ✓ 1494ms | 否 | ✓ 1526ms | ✓ 1307ms | http |
| 103.35.190.69:1081 | ✓ 217ms | ✓ 1190ms | 否 | 否 | ✓ 1753ms | http |
| 45.140.147.82:1081 | ✓ 1012ms | ✓ 1563ms | ✓ 1184ms | ✓ 1787ms | ✓ 1002ms | http |
| 138.68.153.144:3128 | ✓ 1094ms | 否 | ✓ 1999ms | 否 | ✓ 1676ms | http |
| 91.184.241.12:443 | ✓ 815ms | ✓ 1683ms | ✓ 1107ms | ✓ 1993ms | ✓ 1499ms | http |
| 5.253.43.103:3128 | ✓ 1297ms | ✓ 1522ms | ✓ 1469ms | ✓ 1923ms | 否 | http |
| 91.108.243.203:3128 | ✓ 1268ms | ✓ 1920ms | ✓ 1564ms | 否 | 否 | http |
| 103.35.191.138:1082 | ✓ 274ms | ✓ 989ms | ✓ 598ms | ✓ 1315ms | ✓ 989ms | http |
| 121.230.8.55:1080 | 否 | ✓ 1822ms | ✓ 1708ms | ✓ 1953ms | 否 | http |
| 61.52.131.172:8443 | ✓ 1058ms | ✓ 1309ms | ✓ 1105ms | ✓ 1390ms | ✓ 1083ms | http |
| 128.199.121.61:9090 | ✓ 1659ms | 否 | ✓ 1705ms | ✓ 1862ms | 否 | http |
| 45.140.147.155:1082 | ✓ 1107ms | ✓ 1818ms | 否 | 否 | ✓ 1483ms | http |

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
