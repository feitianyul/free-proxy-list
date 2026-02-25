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

最后更新：2026-02-25 16:58:04 UTC（2026-02-26 00:58:04 UTC+8）

**代理总数：60**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 60 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 60 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 85.208.108.43:2094 | 否 | 否 | ✓ 1232ms | ✓ 1514ms | ✓ 807ms | http |
| 104.238.30.50:59741 | ✓ 1581ms | 否 | ✓ 1939ms | 否 | ✓ 1899ms | http |
| 104.238.30.45:59741 | ✓ 1582ms | 否 | ✓ 1935ms | 否 | ✓ 1902ms | http |
| 35.225.22.61:80 | ✓ 891ms | 否 | 否 | ✓ 1068ms | ✓ 867ms | http |
| 20.78.118.91:8561 | ✓ 1573ms | ✓ 1247ms | ✓ 849ms | ✓ 1069ms | ✓ 850ms | http |
| 20.210.39.153:8561 | ✓ 1620ms | ✓ 1698ms | ✓ 689ms | ✓ 1046ms | ✓ 774ms | http |
| 20.78.26.206:8561 | ✓ 1557ms | ✓ 1697ms | ✓ 671ms | ✓ 1054ms | ✓ 775ms | http |
| 103.82.23.118:5234 | ✓ 1647ms | 否 | ✓ 1555ms | ✓ 1982ms | ✓ 1722ms | http |
| 138.124.53.25:7443 | ✓ 1705ms | 否 | ✓ 1860ms | ✓ 1808ms | 否 | http |
| 190.242.157.215:8080 | ✓ 1473ms | 否 | ✓ 1833ms | 否 | ✓ 1572ms | http |
| 202.152.44.19:8081 | ✓ 1413ms | 否 | ✓ 1058ms | ✓ 1486ms | ✓ 1159ms | http |
| 202.152.44.18:8081 | ✓ 1424ms | 否 | ✓ 1152ms | ✓ 1560ms | ✓ 1248ms | http |
| 104.238.30.39:59741 | ✓ 1620ms | 否 | ✓ 1707ms | 否 | ✓ 1935ms | http |
| 72.56.59.23:61937 | ✓ 1442ms | 否 | ✓ 1488ms | 否 | ✓ 1966ms | http |
| 104.238.30.58:63744 | ✓ 1584ms | 否 | ✓ 1903ms | 否 | ✓ 1935ms | http |
| 35.234.17.221:8080 | ✓ 1175ms | ✓ 1563ms | ✓ 1285ms | 否 | ✓ 1025ms | http |
| 72.56.59.56:63127 | ✓ 1441ms | 否 | ✓ 1967ms | 否 | ✓ 1742ms | http |
| 104.238.30.91:63900 | ✓ 1625ms | 否 | ✓ 1935ms | 否 | ✓ 1934ms | http |
| 81.70.169.194:80 | ✓ 1172ms | 否 | ✓ 1157ms | ✓ 1472ms | ✓ 1240ms | http |
| 120.92.212.16:7890 | ✓ 1962ms | 否 | ✓ 1157ms | ✓ 1455ms | 否 | http |
| 195.123.209.48:3128 | ✓ 921ms | 否 | 否 | ✓ 1689ms | ✓ 1693ms | http |
| 45.136.198.40:3128 | ✓ 673ms | ✓ 1881ms | 否 | ✓ 1925ms | 否 | http |
| 138.197.68.35:4857 | 否 | 否 | ✓ 889ms | ✓ 993ms | ✓ 1127ms | http |
| 120.92.212.16:8890 | ✓ 1402ms | 否 | ✓ 1140ms | 否 | ✓ 1165ms | http |
| 91.233.223.147:3128 | ✓ 764ms | 否 | ✓ 735ms | ✓ 1919ms | ✓ 1448ms | http |
| 120.79.99.232:8099 | ✓ 1456ms | ✓ 1793ms | ✓ 1449ms | ✓ 1664ms | ✓ 1392ms | http |
| 115.231.181.40:8128 | ✓ 1050ms | 否 | 否 | ✓ 1470ms | ✓ 1913ms | http |
| 71.204.156.52:443 | ✓ 920ms | ✓ 1562ms | ✓ 1208ms | 否 | 否 | http |
| 104.238.30.37:59741 | ✓ 1986ms | 否 | ✓ 1743ms | 否 | ✓ 1940ms | http |
| 103.82.23.118:5171 | 否 | 否 | ✓ 1673ms | ✓ 1957ms | ✓ 1424ms | http |
| 45.88.0.99:3128 | ✓ 608ms | ✓ 1233ms | 否 | ✓ 1820ms | ✓ 1678ms | http |
| 45.88.0.117:3128 | ✓ 605ms | ✓ 1453ms | 否 | ✓ 1850ms | ✓ 1678ms | http |
| 45.88.0.113:3128 | ✓ 1147ms | ✓ 1525ms | 否 | ✓ 1784ms | ✓ 1667ms | http |
| 45.88.0.98:3128 | ✓ 454ms | 否 | ✓ 1655ms | ✓ 1651ms | ✓ 1674ms | http |
| 45.88.0.115:3128 | ✓ 1031ms | 否 | ✓ 1653ms | ✓ 1653ms | ✓ 1676ms | http |
| 45.88.0.116:3128 | ✓ 996ms | 否 | ✓ 1654ms | ✓ 1663ms | ✓ 1667ms | http |
| 45.88.0.111:3128 | ✓ 564ms | 否 | ✓ 1654ms | ✓ 1653ms | ✓ 1677ms | http |
| 45.88.0.114:3128 | ✓ 996ms | 否 | ✓ 1655ms | ✓ 1666ms | ✓ 1660ms | http |
| 59.46.216.131:30001 | 否 | 否 | ✓ 1442ms | ✓ 1737ms | ✓ 1212ms | http |
| 210.77.18.31:7890 | ✓ 1002ms | ✓ 1584ms | ✓ 1023ms | ✓ 1327ms | ✓ 1053ms | http |
| 104.238.30.40:59741 | ✓ 1616ms | 否 | ✓ 1743ms | 否 | ✓ 1935ms | http |
| 150.249.255.91:3128 | ✓ 1628ms | ✓ 1234ms | ✓ 817ms | 否 | 否 | http |
| 209.38.54.154:8443 | ✓ 851ms | 否 | ✓ 1942ms | ✓ 1584ms | ✓ 1151ms | http |
| 90.84.188.97:8000 | ✓ 1929ms | ✓ 1618ms | 否 | 否 | ✓ 1754ms | http |
| 20.2.83.243:3128 | ✓ 810ms | ✓ 1945ms | 否 | ✓ 1308ms | ✓ 876ms | http |
| 210.223.44.230:3128 | ✓ 819ms | 否 | ✓ 1969ms | ✓ 1468ms | 否 | http |
| 204.48.31.203:80 | ✓ 859ms | 否 | ✓ 1047ms | ✓ 1232ms | ✓ 974ms | http |
| 103.84.95.54:7890 | 否 | 否 | ✓ 848ms | ✓ 1291ms | ✓ 908ms | http |
| 103.149.99.128:3128 | ✓ 1621ms | ✓ 1837ms | 否 | ✓ 1421ms | ✓ 1184ms | http |
| 101.43.255.96:80 | ✓ 1173ms | ✓ 1525ms | ✓ 1230ms | ✓ 1561ms | ✓ 1164ms | http |
| 150.40.179.167:443 | ✓ 1339ms | 否 | ✓ 1283ms | ✓ 1496ms | ✓ 1279ms | http |
| 36.147.78.166:443 | ✓ 1956ms | 否 | ✓ 1892ms | 否 | ✓ 1946ms | http |
| 37.187.109.70:10111 | ✓ 951ms | 否 | ✓ 871ms | ✓ 1626ms | 否 | http |
| 27.254.99.183:8118 | 否 | 否 | ✓ 1315ms | ✓ 1516ms | ✓ 1229ms | http |
| 172.212.68.37:3128 | ✓ 636ms | ✓ 1349ms | 否 | ✓ 1783ms | ✓ 978ms | http |
| 103.39.51.190:8080 | ✓ 1834ms | 否 | 否 | ✓ 1647ms | ✓ 1617ms | http |
| 124.16.93.70:7890 | ✓ 1026ms | ✓ 1596ms | ✓ 1075ms | ✓ 1347ms | ✓ 1081ms | http |
| 103.113.70.189:1081 | 否 | ✓ 964ms | 否 | ✓ 1119ms | ✓ 727ms | http |
| 178.130.47.129:1082 | ✓ 1443ms | 否 | ✓ 1038ms | 否 | ✓ 722ms | http |
| 81.177.48.54:2080 | ✓ 1260ms | 否 | ✓ 1660ms | 否 | ✓ 1918ms | http |

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
