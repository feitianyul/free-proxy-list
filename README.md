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

最后更新：2026-03-01 15:39:40 UTC（2026-03-01 23:39:40 UTC+8）

**代理总数：54**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 54 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 54 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 249ms | ✓ 1050ms | ✓ 873ms | ✓ 1255ms | ✓ 904ms | http |
| 148.135.85.87:1080 | ✓ 839ms | 否 | ✓ 1749ms | 否 | ✓ 1063ms | http |
| 34.101.184.164:3128 | ✓ 1635ms | 否 | ✓ 1454ms | ✓ 1553ms | ✓ 1337ms | http |
| 36.147.78.166:80 | 否 | ✓ 1778ms | ✓ 1772ms | ✓ 1985ms | 否 | http |
| 14.56.107.244:3128 | 否 | ✓ 1464ms | ✓ 1507ms | ✓ 1625ms | 否 | http |
| 35.225.22.61:80 | 否 | 否 | ✓ 696ms | ✓ 971ms | ✓ 823ms | http |
| 59.46.216.131:30001 | ✓ 1868ms | 否 | ✓ 1259ms | ✓ 1519ms | 否 | http |
| 20.78.118.91:8561 | ✓ 1769ms | ✓ 1925ms | ✓ 1703ms | 否 | 否 | http |
| 120.92.212.16:8890 | 否 | ✓ 1564ms | 否 | ✓ 1524ms | ✓ 1018ms | http |
| 141.11.210.35:1080 | 否 | 否 | ✓ 1177ms | ✓ 1220ms | ✓ 706ms | http |
| 45.140.147.155:1082 | ✓ 568ms | ✓ 1818ms | ✓ 1961ms | 否 | 否 | http |
| 20.78.26.206:8561 | 否 | ✓ 1110ms | ✓ 694ms | ✓ 909ms | ✓ 718ms | http |
| 20.210.39.153:8561 | 否 | ✓ 1388ms | ✓ 583ms | ✓ 875ms | ✓ 704ms | http |
| 95.85.252.153:21064 | ✓ 677ms | ✓ 1861ms | ✓ 1383ms | 否 | 否 | http |
| 36.147.78.166:443 | 否 | 否 | ✓ 1803ms | ✓ 1943ms | ✓ 1681ms | http |
| 81.70.169.194:80 | ✓ 1347ms | ✓ 1294ms | ✓ 1092ms | ✓ 1585ms | ✓ 1023ms | http |
| 101.43.255.96:80 | ✓ 1318ms | ✓ 1743ms | ✓ 1026ms | 否 | ✓ 1433ms | http |
| 20.27.15.111:8561 | 否 | 否 | ✓ 692ms | ✓ 1217ms | ✓ 1324ms | http |
| 20.27.14.220:8561 | 否 | 否 | ✓ 613ms | ✓ 1203ms | ✓ 1354ms | http |
| 20.27.11.248:8561 | 否 | 否 | ✓ 802ms | ✓ 1436ms | ✓ 1191ms | http |
| 165.227.5.10:8888 | ✓ 375ms | 否 | 否 | ✓ 1829ms | ✓ 1003ms | http |
| 142.171.85.32:1080 | ✓ 824ms | ✓ 1961ms | 否 | ✓ 856ms | 否 | http |
| 121.128.121.184:3128 | 否 | 否 | ✓ 1996ms | ✓ 1430ms | ✓ 1942ms | http |
| 47.101.149.27:9010 | 否 | 否 | ✓ 1888ms | ✓ 1565ms | ✓ 1977ms | http |
| 62.113.119.14:8080 | ✓ 1081ms | ✓ 1606ms | ✓ 598ms | ✓ 1560ms | ✓ 1134ms | http |
| 168.235.110.63:3128 | ✓ 191ms | ✓ 1752ms | ✓ 774ms | 否 | 否 | http |
| 45.140.147.155:1081 | ✓ 486ms | 否 | ✓ 1467ms | ✓ 1778ms | ✓ 1581ms | http |
| 115.231.181.40:8128 | 否 | ✓ 1142ms | ✓ 1997ms | 否 | ✓ 1733ms | http |
| 103.104.99.29:80 | 否 | 否 | ✓ 1720ms | ✓ 1630ms | ✓ 1599ms | http |
| 103.104.99.89:80 | 否 | 否 | ✓ 1560ms | ✓ 1637ms | ✓ 1651ms | http |
| 61.72.221.234:3128 | ✓ 1633ms | 否 | ✓ 966ms | 否 | ✓ 1948ms | http |
| 5.75.201.136:1080 | ✓ 816ms | 否 | 否 | ✓ 1705ms | ✓ 1445ms | http |
| 198.244.151.77:8888 | ✓ 1055ms | 否 | ✓ 1367ms | 否 | ✓ 1990ms | http |
| 138.124.53.25:7443 | 否 | 否 | ✓ 1661ms | ✓ 1676ms | ✓ 1682ms | http |
| 118.31.1.154:80 | ✓ 917ms | ✓ 1195ms | ✓ 998ms | ✓ 1208ms | ✓ 974ms | http |
| 35.234.17.221:8080 | 否 | 否 | ✓ 864ms | ✓ 1346ms | ✓ 1003ms | http |
| 120.92.212.16:7890 | ✓ 1062ms | 否 | ✓ 1311ms | 否 | ✓ 1039ms | http |
| 45.136.198.40:3128 | ✓ 739ms | 否 | ✓ 1778ms | 否 | ✓ 1710ms | http |
| 20.210.76.104:8561 | ✓ 1205ms | ✓ 1371ms | ✓ 1465ms | 否 | 否 | http |
| 20.210.76.175:8561 | ✓ 1206ms | ✓ 1371ms | ✓ 1465ms | 否 | 否 | http |
| 45.125.67.37:8443 | ✓ 1293ms | 否 | ✓ 1130ms | 否 | ✓ 1274ms | http |
| 45.177.178.23:999 | ✓ 262ms | ✓ 1087ms | ✓ 286ms | ✓ 1040ms | ✓ 895ms | http |
| 150.107.140.238:3128 | ✓ 1755ms | 否 | 否 | ✓ 1846ms | ✓ 963ms | http |
| 217.119.129.86:2222 | ✓ 1319ms | 否 | ✓ 1072ms | ✓ 1917ms | 否 | http |
| 160.250.4.245:1 | ✓ 1096ms | 否 | ✓ 1167ms | ✓ 1367ms | 否 | http |
| 196.70.95.87:3128 | ✓ 1614ms | 否 | 否 | ✓ 1884ms | ✓ 1812ms | http |
| 103.39.51.190:8080 | ✓ 1986ms | 否 | 否 | ✓ 1398ms | ✓ 1521ms | http |
| 139.159.99.242:8080 | 否 | 否 | ✓ 909ms | ✓ 1157ms | ✓ 893ms | http |
| 121.128.121.54:3128 | ✓ 1179ms | ✓ 1666ms | 否 | 否 | ✓ 923ms | http |
| 14.56.177.44:3128 | ✓ 1168ms | ✓ 1755ms | ✓ 1578ms | 否 | 否 | http |
| 107.174.133.10:3128 | ✓ 903ms | 否 | 否 | ✓ 1257ms | ✓ 993ms | http |
| 172.212.68.37:3128 | ✓ 338ms | ✓ 1670ms | 否 | ✓ 1544ms | ✓ 823ms | http |
| 103.84.95.54:7890 | ✓ 1162ms | 否 | 否 | ✓ 1250ms | ✓ 967ms | http |
| 157.245.194.13:8888 | ✓ 1575ms | 否 | ✓ 1539ms | ✓ 1157ms | ✓ 1353ms | http |

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
