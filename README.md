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

最后更新：2026-03-10 06:44:19 UTC（2026-03-10 14:44:19 UTC+8）

**代理总数：66**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 65 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 66 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 1.231.81.166:3128 | ✓ 772ms | ✓ 1098ms | ✓ 1050ms | ✓ 1044ms | ✓ 898ms | http |
| 154.3.236.202:3128 | ✓ 500ms | 否 | ✓ 1160ms | ✓ 1619ms | ✓ 1172ms | http |
| 35.225.22.61:80 | ✓ 663ms | ✓ 1273ms | ✓ 265ms | 否 | 否 | http |
| 45.140.147.82:1081 | ✓ 540ms | ✓ 1455ms | ✓ 409ms | ✓ 1720ms | ✓ 1203ms | http |
| 178.236.245.59:3128 | ✓ 678ms | 否 | ✓ 1532ms | ✓ 1801ms | ✓ 1669ms | http |
| 178.236.245.17:3128 | ✓ 716ms | 否 | ✓ 1533ms | ✓ 1850ms | ✓ 1630ms | http |
| 46.39.105.157:8080 | ✓ 1673ms | 否 | ✓ 1667ms | ✓ 1612ms | ✓ 1290ms | http |
| 45.136.130.223:8443 | ✓ 521ms | ✓ 1132ms | ✓ 281ms | ✓ 899ms | ✓ 711ms | http |
| 45.136.131.47:8443 | ✓ 285ms | ✓ 841ms | ✓ 238ms | ✓ 1799ms | ✓ 658ms | http |
| 136.49.34.18:8888 | ✓ 1597ms | 否 | ✓ 1393ms | ✓ 1757ms | 否 | http |
| 103.139.138.194:3128 | 否 | 否 | ✓ 1624ms | ✓ 1789ms | ✓ 1237ms | http |
| 168.235.110.63:3128 | ✓ 794ms | ✓ 1059ms | ✓ 1553ms | 否 | 否 | http |
| 103.166.185.54:3128 | ✓ 1882ms | ✓ 1761ms | 否 | ✓ 1348ms | ✓ 1159ms | http |
| 194.213.18.200:443 | ✓ 949ms | ✓ 1654ms | 否 | ✓ 1909ms | 否 | http |
| 91.107.141.42:8081 | ✓ 957ms | ✓ 1722ms | ✓ 1408ms | ✓ 1884ms | ✓ 1526ms | http |
| 193.168.173.136:443 | ✓ 1038ms | 否 | ✓ 1035ms | ✓ 1571ms | ✓ 1386ms | http |
| 39.104.201.40:7890 | ✓ 1063ms | ✓ 1367ms | ✓ 1131ms | ✓ 1368ms | ✓ 1085ms | http |
| 115.231.181.40:8128 | ✓ 1793ms | ✓ 1341ms | ✓ 1074ms | ✓ 1536ms | ✓ 1107ms | http |
| 121.237.181.137:8888 | ✓ 1031ms | 否 | ✓ 1030ms | ✓ 1621ms | ✓ 1743ms | http |
| 120.92.212.16:8890 | 否 | 否 | ✓ 1833ms | ✓ 1412ms | ✓ 1095ms | http |
| 81.70.169.194:80 | ✓ 1149ms | ✓ 1413ms | ✓ 1746ms | ✓ 1449ms | ✓ 1145ms | http |
| 101.43.255.96:80 | ✓ 1744ms | 否 | ✓ 1198ms | ✓ 1332ms | ✓ 1061ms | http |
| 14.225.212.37:7890 | ✓ 1857ms | ✓ 1564ms | ✓ 968ms | ✓ 1294ms | 否 | http |
| 14.225.222.213:7890 | ✓ 1404ms | ✓ 1832ms | ✓ 1783ms | 否 | 否 | http |
| 185.191.236.162:3128 | ✓ 946ms | 否 | ✓ 1302ms | ✓ 1827ms | ✓ 1440ms | http |
| 104.248.195.47:8080 | ✓ 929ms | 否 | 否 | ✓ 1643ms | ✓ 1335ms | http |
| 43.225.185.4:8000 | ✓ 1936ms | 否 | ✓ 1313ms | ✓ 1480ms | ✓ 1161ms | http |
| 45.140.147.155:1082 | ✓ 457ms | 否 | ✓ 970ms | ✓ 1861ms | ✓ 1214ms | http |
| 202.155.12.161:443 | ✓ 1642ms | 否 | ✓ 1368ms | 否 | ✓ 1279ms | http |
| 8.137.112.117:3128 | ✓ 1150ms | ✓ 1439ms | ✓ 1176ms | ✓ 1533ms | ✓ 1130ms | http |
| 95.3.9.78:3128 | ✓ 1900ms | ✓ 1974ms | 否 | ✓ 1886ms | ✓ 1389ms | http |
| 103.35.188.243:3128 | ✓ 1523ms | ✓ 1058ms | 否 | ✓ 1099ms | ✓ 870ms | http |
| 107.172.125.217:3128 | ✓ 279ms | 否 | ✓ 1828ms | ✓ 1086ms | ✓ 827ms | http |
| 46.183.25.8:443 | ✓ 1282ms | 否 | ✓ 828ms | ✓ 1073ms | 否 | http |
| 62.113.119.14:8080 | ✓ 1567ms | 否 | ✓ 1849ms | ✓ 1742ms | ✓ 1831ms | http |
| 210.223.44.230:3128 | ✓ 1043ms | ✓ 1546ms | ✓ 806ms | ✓ 1374ms | ✓ 900ms | http |
| 192.73.243.98:3128 | ✓ 262ms | ✓ 1799ms | 否 | ✓ 1414ms | ✓ 1925ms | http |
| 138.124.53.25:7443 | 否 | 否 | ✓ 1352ms | ✓ 1532ms | ✓ 1581ms | http |
| 116.80.63.46:7777 | ✓ 1617ms | 否 | 否 | ✓ 1979ms | ✓ 1765ms | http |
| 152.70.98.46:8888 | ✓ 1697ms | 否 | ✓ 1880ms | ✓ 1204ms | ✓ 1708ms | http |
| 88.80.150.82:8080 | ✓ 1024ms | ✓ 1897ms | 否 | 否 | ✓ 1711ms | https |
| 121.230.9.198:1080 | 否 | 否 | ✓ 1543ms | ✓ 1760ms | ✓ 1305ms | http |
| 190.9.109.198:999 | 否 | 否 | ✓ 1759ms | ✓ 1292ms | ✓ 1157ms | http |
| 45.136.198.40:3128 | ✓ 1066ms | 否 | ✓ 1769ms | ✓ 1939ms | ✓ 1680ms | http |
| 120.92.212.16:7890 | ✓ 1109ms | ✓ 1375ms | 否 | 否 | ✓ 1064ms | http |
| 103.113.70.189:1081 | ✓ 292ms | ✓ 1592ms | 否 | ✓ 1411ms | ✓ 1552ms | http |
| 190.212.131.238:3128 | ✓ 914ms | 否 | ✓ 1916ms | 否 | ✓ 1625ms | http |
| 43.167.227.161:1080 | ✓ 1800ms | 否 | 否 | ✓ 1779ms | ✓ 783ms | http |
| 185.92.69.160:3128 | ✓ 1013ms | ✓ 1461ms | ✓ 1326ms | 否 | ✓ 1294ms | http |
| 150.249.255.91:3128 | ✓ 1765ms | 否 | ✓ 731ms | ✓ 1051ms | ✓ 1046ms | http |
| 91.233.223.147:3128 | ✓ 869ms | 否 | ✓ 1768ms | ✓ 1984ms | ✓ 1512ms | http |
| 152.69.229.220:3128 | ✓ 1726ms | 否 | 否 | ✓ 1571ms | ✓ 1184ms | http |
| 101.47.73.135:3128 | ✓ 1899ms | 否 | 否 | ✓ 1554ms | ✓ 1773ms | http |
| 162.248.165.72:1080 | ✓ 1669ms | 否 | ✓ 1994ms | 否 | ✓ 1787ms | http |
| 20.120.225.109:3128 | ✓ 400ms | 否 | ✓ 895ms | ✓ 1409ms | 否 | http |
| 106.14.203.63:3333 | ✓ 1973ms | ✓ 1819ms | ✓ 1070ms | 否 | 否 | http |
| 103.39.51.190:8080 | ✓ 1848ms | 否 | 否 | ✓ 1624ms | ✓ 1580ms | http |
| 94.176.3.43:7443 | 否 | 否 | ✓ 1462ms | ✓ 1837ms | ✓ 1570ms | http |
| 45.22.209.157:8888 | ✓ 765ms | 否 | ✓ 1550ms | ✓ 1598ms | 否 | http |
| 61.52.131.172:8443 | ✓ 1107ms | ✓ 1405ms | ✓ 1048ms | ✓ 1313ms | ✓ 1037ms | http |
| 210.77.29.245:7890 | ✓ 1037ms | ✓ 1292ms | ✓ 1130ms | ✓ 1297ms | ✓ 1085ms | http |
| 147.45.251.242:8888 | ✓ 1405ms | 否 | ✓ 1791ms | 否 | ✓ 1872ms | http |
| 205.209.118.30:3138 | ✓ 201ms | ✓ 1305ms | ✓ 856ms | ✓ 1137ms | ✓ 882ms | http |
| 150.107.140.238:3128 | ✓ 1718ms | 否 | ✓ 1202ms | ✓ 1337ms | 否 | http |
| 59.46.216.131:30001 | ✓ 1432ms | 否 | ✓ 1199ms | ✓ 1523ms | 否 | http |
| 34.101.184.164:3128 | ✓ 1598ms | 否 | ✓ 1430ms | ✓ 1417ms | ✓ 1483ms | http |

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
