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

最后更新：2026-03-03 00:24:56 UTC（2026-03-03 08:24:56 UTC+8）

**代理总数：56**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 56 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 56 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 166.0.192.117:8888 | ✓ 1001ms | ✓ 1806ms | ✓ 1795ms | ✓ 890ms | ✓ 676ms | http |
| 91.99.99.83:9000 | ✓ 602ms | 否 | 否 | ✓ 1961ms | ✓ 1646ms | http |
| 45.140.147.82:1081 | ✓ 520ms | 否 | 否 | ✓ 1519ms | ✓ 1228ms | http |
| 186.148.180.46:999 | 否 | ✓ 1797ms | ✓ 1365ms | ✓ 1869ms | ✓ 1396ms | http |
| 35.225.22.61:80 | ✓ 369ms | 否 | ✓ 310ms | ✓ 1283ms | 否 | http |
| 205.209.118.30:3138 | ✓ 293ms | 否 | 否 | ✓ 1265ms | ✓ 978ms | http |
| 223.16.170.103:80 | 否 | 否 | ✓ 1094ms | ✓ 1165ms | ✓ 1325ms | http |
| 210.223.44.230:3128 | ✓ 1164ms | ✓ 1191ms | 否 | ✓ 1042ms | ✓ 1097ms | http |
| 14.56.177.44:3128 | ✓ 1332ms | 否 | ✓ 1002ms | ✓ 1061ms | 否 | http |
| 61.72.110.54:3128 | ✓ 1787ms | ✓ 1715ms | ✓ 1540ms | 否 | 否 | http |
| 101.43.255.96:80 | ✓ 947ms | ✓ 1238ms | ✓ 1086ms | ✓ 1337ms | ✓ 985ms | http |
| 81.70.169.194:80 | ✓ 1192ms | ✓ 1346ms | ✓ 1261ms | ✓ 1330ms | ✓ 970ms | http |
| 195.123.209.48:3128 | ✓ 1173ms | ✓ 1785ms | ✓ 1518ms | 否 | ✓ 1818ms | http |
| 118.113.247.73:1080 | ✓ 1397ms | ✓ 1733ms | ✓ 1562ms | 否 | ✓ 1896ms | http |
| 120.92.212.16:7890 | ✓ 1012ms | 否 | 否 | ✓ 1454ms | ✓ 1154ms | http |
| 115.76.5.32:10006 | 否 | 否 | ✓ 1585ms | ✓ 1586ms | ✓ 1537ms | http |
| 103.84.95.54:7890 | ✓ 1050ms | 否 | ✓ 765ms | 否 | ✓ 1591ms | http |
| 114.231.73.6:1080 | ✓ 979ms | ✓ 1319ms | ✓ 988ms | ✓ 1292ms | ✓ 956ms | http |
| 14.103.233.245:3128 | ✓ 1368ms | ✓ 1136ms | ✓ 932ms | ✓ 1103ms | ✓ 908ms | http |
| 115.231.181.40:8128 | 否 | ✓ 1066ms | 否 | ✓ 1175ms | ✓ 973ms | http |
| 34.101.184.164:3128 | 否 | 否 | ✓ 1057ms | ✓ 1462ms | ✓ 1056ms | http |
| 35.234.17.221:8080 | ✓ 909ms | ✓ 1711ms | ✓ 1422ms | ✓ 1562ms | ✓ 916ms | http |
| 115.76.5.32:10009 | 否 | 否 | ✓ 1653ms | ✓ 1836ms | ✓ 1513ms | http |
| 120.92.212.16:8890 | 否 | 否 | ✓ 1222ms | ✓ 1494ms | ✓ 1190ms | http |
| 5.75.196.26:40000 | ✓ 544ms | ✓ 1735ms | ✓ 1499ms | 否 | 否 | http |
| 115.76.5.32:10007 | 否 | 否 | ✓ 1462ms | ✓ 1945ms | ✓ 1796ms | http |
| 115.76.5.32:10008 | ✓ 1897ms | 否 | ✓ 1422ms | 否 | ✓ 1547ms | http |
| 180.127.149.247:1080 | ✓ 980ms | ✓ 1199ms | ✓ 1022ms | ✓ 1259ms | ✓ 913ms | http |
| 121.230.9.161:1080 | ✓ 1178ms | ✓ 1446ms | ✓ 1059ms | ✓ 1742ms | ✓ 1368ms | http |
| 62.113.119.14:8080 | ✓ 1222ms | 否 | ✓ 1102ms | ✓ 1571ms | ✓ 1195ms | http |
| 142.171.85.32:1080 | 否 | 否 | ✓ 1655ms | ✓ 1717ms | ✓ 1965ms | http |
| 181.78.49.177:999 | ✓ 1756ms | 否 | ✓ 993ms | ✓ 1844ms | 否 | http |
| 45.125.67.37:8443 | ✓ 934ms | 否 | ✓ 987ms | 否 | ✓ 997ms | http |
| 160.238.65.8:3128 | ✓ 1743ms | 否 | ✓ 830ms | ✓ 1757ms | 否 | http |
| 3.213.157.4:3128 | ✓ 356ms | ✓ 1738ms | 否 | ✓ 1750ms | ✓ 1217ms | http |
| 138.124.53.25:7443 | ✓ 1241ms | 否 | ✓ 1410ms | ✓ 1885ms | ✓ 1505ms | http |
| 103.35.188.243:3128 | 否 | ✓ 1253ms | 否 | ✓ 1305ms | ✓ 1115ms | http |
| 46.249.103.192:443 | ✓ 1242ms | 否 | ✓ 1305ms | ✓ 1637ms | 否 | http |
| 121.128.121.54:3128 | ✓ 1290ms | 否 | ✓ 903ms | 否 | ✓ 911ms | http |
| 38.180.2.107:3128 | ✓ 874ms | 否 | ✓ 1773ms | 否 | ✓ 1752ms | http |
| 45.140.147.82:1082 | ✓ 614ms | ✓ 1917ms | ✓ 1342ms | ✓ 1292ms | ✓ 1396ms | http |
| 188.166.208.168:9876 | ✓ 1018ms | 否 | ✓ 743ms | ✓ 1087ms | ✓ 843ms | http |
| 45.136.198.40:3128 | ✓ 1106ms | ✓ 1955ms | 否 | 否 | ✓ 1892ms | http |
| 95.85.252.153:21064 | ✓ 536ms | ✓ 1559ms | ✓ 997ms | 否 | ✓ 1589ms | http |
| 103.39.51.190:8080 | ✓ 1746ms | 否 | 否 | ✓ 1400ms | ✓ 1340ms | http |
| 103.215.36.88:15556 | ✓ 1181ms | ✓ 1466ms | ✓ 1201ms | ✓ 1411ms | ✓ 1105ms | http |
| 158.160.215.167:8127 | ✓ 1277ms | 否 | ✓ 1118ms | 否 | ✓ 1562ms | http |
| 122.2.48.121:8080 | 否 | 否 | ✓ 1431ms | ✓ 1313ms | ✓ 1294ms | http |
| 91.233.223.147:3128 | ✓ 1058ms | 否 | ✓ 1098ms | 否 | ✓ 1777ms | http |
| 61.72.221.194:3128 | 否 | 否 | ✓ 1909ms | ✓ 1366ms | ✓ 1798ms | http |
| 103.215.36.88:15852 | ✓ 1283ms | ✓ 1395ms | ✓ 1285ms | ✓ 1422ms | ✓ 1124ms | http |
| 137.184.14.135:3128 | 否 | ✓ 1560ms | ✓ 1544ms | ✓ 1564ms | ✓ 1196ms | http |
| 114.231.72.60:1080 | ✓ 973ms | ✓ 1221ms | ✓ 1016ms | ✓ 1224ms | ✓ 875ms | http |
| 91.238.104.171:2023 | ✓ 985ms | 否 | ✓ 1014ms | 否 | ✓ 1265ms | http |
| 193.123.60.230:3128 | ✓ 1416ms | 否 | 否 | ✓ 1956ms | ✓ 1768ms | http |
| 103.215.36.88:10864 | ✓ 1170ms | ✓ 1302ms | ✓ 1077ms | ✓ 1316ms | ✓ 1052ms | http |

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
