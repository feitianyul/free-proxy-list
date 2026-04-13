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

最后更新：2026-04-13 18:05:14 UTC（2026-04-14 02:05:14 UTC+8）

**代理总数：67**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 67 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 67 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 147.161.210.140:8800 | ✓ 1633ms | 否 | ✓ 1522ms | ✓ 1247ms | ✓ 1095ms | http |
| 167.103.115.102:8800 | ✓ 1594ms | 否 | ✓ 1134ms | ✓ 1821ms | ✓ 1236ms | http |
| 167.103.34.108:8800 | ✓ 1649ms | 否 | ✓ 1335ms | ✓ 1779ms | 否 | http |
| 43.156.132.113:3128 | ✓ 1593ms | ✓ 1647ms | 否 | ✓ 1748ms | 否 | http |
| 8.219.195.129:1080 | ✓ 973ms | ✓ 1833ms | ✓ 834ms | ✓ 1187ms | ✓ 986ms | http |
| 167.103.144.127:8800 | ✓ 1539ms | 否 | ✓ 1770ms | 否 | ✓ 1565ms | http |
| 162.240.154.26:3128 | ✓ 1135ms | 否 | ✓ 1720ms | ✓ 1005ms | 否 | http |
| 1.231.81.166:3128 | ✓ 1528ms | ✓ 1488ms | ✓ 1971ms | ✓ 1818ms | 否 | http |
| 43.99.54.236:5555 | ✓ 1892ms | 否 | 否 | ✓ 1444ms | ✓ 769ms | http |
| 45.167.125.21:999 | ✓ 1151ms | 否 | ✓ 765ms | ✓ 1910ms | ✓ 1446ms | http |
| 137.59.47.73:3128 | ✓ 1545ms | ✓ 1537ms | ✓ 1008ms | ✓ 1209ms | ✓ 1040ms | http |
| 46.30.46.133:3128 | ✓ 736ms | 否 | ✓ 502ms | ✓ 1491ms | ✓ 1492ms | http |
| 171.227.167.109:1005 | ✓ 1849ms | 否 | ✓ 1562ms | ✓ 1761ms | ✓ 1194ms | http |
| 207.254.71.62:8088 | ✓ 889ms | ✓ 1698ms | ✓ 1278ms | ✓ 1899ms | ✓ 1875ms | http |
| 45.140.147.155:1082 | ✓ 815ms | 否 | ✓ 468ms | ✓ 1737ms | ✓ 971ms | http |
| 45.140.147.155:1081 | ✓ 808ms | 否 | ✓ 424ms | ✓ 1789ms | ✓ 985ms | http |
| 167.103.31.122:8800 | ✓ 1412ms | 否 | ✓ 1384ms | ✓ 1975ms | ✓ 1643ms | http |
| 218.108.131.186:17890 | ✓ 1419ms | ✓ 1994ms | ✓ 1349ms | 否 | 否 | http |
| 159.223.225.118:8888 | ✓ 621ms | 否 | ✓ 1444ms | ✓ 1839ms | ✓ 1251ms | http |
| 34.71.229.255:3128 | ✓ 1441ms | 否 | ✓ 1145ms | ✓ 1157ms | ✓ 1241ms | http |
| 147.161.239.240:8800 | ✓ 999ms | ✓ 1622ms | ✓ 1726ms | ✓ 1660ms | ✓ 1338ms | http |
| 101.43.127.100:8877 | ✓ 1057ms | ✓ 1293ms | ✓ 1006ms | ✓ 1341ms | ✓ 1089ms | http |
| 113.160.132.26:8080 | ✓ 1980ms | 否 | 否 | ✓ 1436ms | ✓ 1126ms | http |
| 185.191.236.162:3128 | ✓ 1099ms | 否 | ✓ 1641ms | 否 | ✓ 1530ms | http |
| 123.57.2.231:2020 | ✓ 894ms | ✓ 1118ms | ✓ 876ms | ✓ 1146ms | ✓ 929ms | http |
| 130.61.30.221:8080 | ✓ 1020ms | 否 | ✓ 1593ms | ✓ 1966ms | ✓ 1442ms | http |
| 152.32.132.190:7890 | ✓ 1350ms | 否 | ✓ 953ms | 否 | ✓ 1280ms | http |
| 212.58.132.5:8888 | ✓ 1077ms | 否 | ✓ 1308ms | ✓ 1625ms | ✓ 1563ms | http |
| 54.37.72.89:80 | ✓ 1321ms | 否 | ✓ 1909ms | 否 | ✓ 1933ms | http |
| 36.141.21.200:7890 | ✓ 1031ms | ✓ 1368ms | 否 | 否 | ✓ 1221ms | http |
| 147.45.167.84:3128 | 否 | 否 | ✓ 1690ms | ✓ 1684ms | ✓ 1716ms | http |
| 152.32.93.105:8082 | ✓ 1486ms | 否 | 否 | ✓ 1476ms | ✓ 1490ms | http |
| 5.104.87.17:8051 | ✓ 1524ms | 否 | ✓ 1324ms | ✓ 1616ms | ✓ 1409ms | http |
| 35.225.22.61:80 | 否 | 否 | ✓ 281ms | ✓ 1067ms | ✓ 872ms | http |
| 34.101.184.164:3128 | ✓ 919ms | 否 | ✓ 1067ms | ✓ 1398ms | ✓ 1518ms | http |
| 210.223.44.230:3128 | 否 | ✓ 1403ms | ✓ 1649ms | 否 | ✓ 1332ms | http |
| 120.92.108.86:7890 | ✓ 1428ms | 否 | ✓ 1343ms | ✓ 1823ms | ✓ 1448ms | http |
| 59.46.216.131:30001 | ✓ 1410ms | ✓ 1576ms | ✓ 1186ms | 否 | 否 | http |
| 116.203.112.97:3128 | ✓ 518ms | ✓ 1549ms | 否 | 否 | ✓ 1913ms | http |
| 119.93.151.167:5050 | ✓ 1538ms | 否 | ✓ 1980ms | ✓ 1494ms | 否 | http |
| 121.230.8.133:1080 | ✓ 1206ms | 否 | ✓ 1200ms | ✓ 1438ms | 否 | http |
| 121.230.8.235:1080 | ✓ 1341ms | 否 | ✓ 1316ms | ✓ 1776ms | ✓ 1204ms | http |
| 163.61.38.126:3128 | ✓ 1828ms | 否 | 否 | ✓ 1950ms | ✓ 1541ms | http |
| 163.61.38.128:3128 | ✓ 1826ms | 否 | 否 | ✓ 1895ms | ✓ 1566ms | http |
| 163.61.38.127:3128 | ✓ 1847ms | 否 | 否 | ✓ 1869ms | ✓ 1655ms | http |
| 168.222.254.88:3128 | ✓ 1159ms | 否 | ✓ 978ms | ✓ 1792ms | ✓ 1759ms | http |
| 181.78.44.63:999 | ✓ 724ms | 否 | ✓ 1015ms | ✓ 1700ms | ✓ 1386ms | http |
| 91.233.223.147:3128 | ✓ 1465ms | 否 | ✓ 1636ms | 否 | ✓ 1788ms | http |
| 36.103.198.235:7890 | 否 | ✓ 1652ms | 否 | ✓ 1538ms | ✓ 1493ms | http |
| 114.237.77.199:1080 | 否 | ✓ 1245ms | ✓ 1819ms | 否 | ✓ 1352ms | http |
| 79.132.136.58:3128 | ✓ 1094ms | ✓ 1486ms | ✓ 1524ms | 否 | 否 | http |
| 113.59.113.4:1088 | ✓ 884ms | 否 | ✓ 1030ms | 否 | ✓ 915ms | http |
| 139.159.99.242:8080 | ✓ 946ms | ✓ 1134ms | ✓ 951ms | ✓ 1195ms | ✓ 986ms | http |
| 154.17.27.171:3128 | ✓ 465ms | ✓ 938ms | ✓ 708ms | ✓ 994ms | ✓ 781ms | http |
| 112.31.254.207:7897 | ✓ 845ms | ✓ 1027ms | ✓ 928ms | ✓ 1502ms | ✓ 877ms | http |
| 185.76.240.61:10001 | ✓ 1126ms | 否 | ✓ 1144ms | ✓ 1955ms | ✓ 1704ms | http |
| 185.76.240.203:10001 | ✓ 1663ms | 否 | ✓ 1022ms | ✓ 1966ms | 否 | http |
| 119.195.17.15:3068 | 否 | 否 | ✓ 1216ms | ✓ 1635ms | ✓ 1174ms | http |
| 94.131.118.129:1081 | ✓ 648ms | ✓ 1228ms | ✓ 671ms | ✓ 1607ms | ✓ 1218ms | http |
| 223.84.151.86:30005 | ✓ 1359ms | ✓ 1649ms | ✓ 1608ms | 否 | 否 | http |
| 190.14.231.47:999 | ✓ 1141ms | 否 | ✓ 1363ms | ✓ 1682ms | ✓ 1930ms | http |
| 113.176.92.71:3128 | ✓ 1462ms | ✓ 1445ms | ✓ 1040ms | ✓ 1327ms | 否 | http |
| 47.105.98.23:3128 | ✓ 1107ms | 否 | 否 | ✓ 1128ms | ✓ 1461ms | http |
| 103.67.46.225:3125 | 否 | 否 | ✓ 1848ms | ✓ 1795ms | ✓ 1683ms | http |
| 94.131.118.39:1081 | ✓ 1932ms | 否 | ✓ 1459ms | ✓ 1493ms | 否 | http |
| 103.18.78.250:1111 | 否 | 否 | ✓ 1334ms | ✓ 1433ms | ✓ 1738ms | http |
| 173.212.246.157:3128 | ✓ 1126ms | 否 | ✓ 1213ms | 否 | ✓ 1850ms | http |

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
