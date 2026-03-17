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

最后更新：2026-03-17 10:57:05 UTC（2026-03-17 18:57:05 UTC+8）

**代理总数：42**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 41 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 42 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 202.155.12.161:443 | ✓ 1689ms | 否 | ✓ 926ms | ✓ 1217ms | ✓ 1180ms | http |
| 147.161.210.140:8800 | ✓ 1690ms | 否 | ✓ 798ms | ✓ 1445ms | ✓ 1127ms | http |
| 219.117.204.211:7799 | ✓ 1909ms | ✓ 1671ms | ✓ 1164ms | ✓ 1822ms | ✓ 1033ms | http |
| 45.167.124.52:8080 | ✓ 1626ms | 否 | ✓ 1257ms | ✓ 1689ms | ✓ 1440ms | http |
| 101.47.73.135:3128 | ✓ 1537ms | 否 | ✓ 1707ms | 否 | ✓ 1316ms | http |
| 137.220.150.170:6005 | ✓ 1558ms | 否 | ✓ 889ms | ✓ 1292ms | ✓ 1365ms | http |
| 113.160.132.26:8080 | ✓ 1647ms | 否 | ✓ 1166ms | ✓ 1467ms | ✓ 1277ms | http |
| 38.145.208.161:8444 | ✓ 1209ms | ✓ 1097ms | ✓ 583ms | ✓ 1453ms | ✓ 855ms | http |
| 38.145.208.182:8453 | ✓ 864ms | ✓ 1678ms | ✓ 356ms | ✓ 1914ms | ✓ 841ms | http |
| 38.145.203.87:8453 | ✓ 821ms | ✓ 1578ms | 否 | ✓ 1056ms | ✓ 1116ms | http |
| 62.60.177.204:34094 | ✓ 1089ms | ✓ 1138ms | 否 | ✓ 1051ms | ✓ 947ms | http |
| 1.231.81.166:3128 | ✓ 1156ms | ✓ 1341ms | ✓ 1733ms | ✓ 1225ms | ✓ 1280ms | http |
| 45.136.198.40:3128 | ✓ 1116ms | 否 | ✓ 1402ms | 否 | ✓ 1431ms | http |
| 65.108.203.36:18080 | ✓ 1509ms | ✓ 1813ms | ✓ 928ms | 否 | 否 | http |
| 115.231.181.40:8128 | ✓ 1032ms | 否 | ✓ 1114ms | ✓ 1250ms | 否 | http |
| 165.227.5.10:8888 | ✓ 1758ms | 否 | 否 | ✓ 1133ms | ✓ 1867ms | http |
| 185.115.74.185:8080 | ✓ 1611ms | ✓ 1676ms | ✓ 1823ms | 否 | 否 | http |
| 137.220.151.110:6005 | ✓ 1516ms | 否 | ✓ 1351ms | ✓ 1223ms | ✓ 974ms | http |
| 35.225.22.61:80 | 否 | ✓ 1710ms | 否 | ✓ 1063ms | ✓ 858ms | http |
| 8.219.97.248:80 | 否 | 否 | ✓ 1271ms | ✓ 1663ms | ✓ 1361ms | http |
| 120.92.212.16:7890 | ✓ 1065ms | ✓ 1392ms | 否 | ✓ 1411ms | ✓ 1358ms | http |
| 193.23.200.251:10808 | ✓ 1269ms | 否 | ✓ 1871ms | ✓ 1753ms | ✓ 1768ms | http |
| 137.220.150.152:6005 | ✓ 1681ms | 否 | ✓ 918ms | ✓ 1392ms | ✓ 1022ms | http |
| 88.80.150.82:8080 | ✓ 1088ms | 否 | ✓ 1859ms | 否 | ✓ 1642ms | https |
| 137.220.150.104:6005 | ✓ 1607ms | 否 | ✓ 865ms | ✓ 1629ms | ✓ 1379ms | http |
| 120.92.212.16:8890 | ✓ 1103ms | ✓ 1402ms | 否 | ✓ 1371ms | 否 | http |
| 192.71.213.85:9091 | ✓ 1502ms | 否 | ✓ 1993ms | ✓ 1829ms | 否 | http |
| 45.88.0.111:3128 | 否 | 否 | ✓ 1968ms | ✓ 1820ms | ✓ 1793ms | http |
| 106.117.208.101:7890 | ✓ 1155ms | ✓ 1464ms | 否 | 否 | ✓ 1171ms | http |
| 149.50.116.240:1080 | ✓ 518ms | 否 | ✓ 1650ms | ✓ 1771ms | 否 | http |
| 45.140.147.155:1082 | ✓ 833ms | 否 | ✓ 1262ms | ✓ 1675ms | ✓ 1226ms | http |
| 45.140.147.155:1081 | 否 | 否 | ✓ 1115ms | ✓ 1609ms | ✓ 1174ms | http |
| 138.124.53.25:7443 | 否 | 否 | ✓ 1271ms | ✓ 1581ms | ✓ 1355ms | http |
| 52.74.26.202:8080 | ✓ 1206ms | ✓ 1666ms | ✓ 889ms | ✓ 1197ms | ✓ 936ms | http |
| 101.43.127.100:8877 | 否 | ✓ 1213ms | ✓ 1066ms | 否 | ✓ 1005ms | http |
| 106.14.203.63:3333 | ✓ 985ms | ✓ 1205ms | ✓ 970ms | ✓ 1240ms | ✓ 982ms | http |
| 103.39.51.190:8080 | ✓ 1900ms | 否 | 否 | ✓ 1815ms | ✓ 1543ms | http |
| 85.198.96.242:3128 | 否 | 否 | ✓ 565ms | ✓ 1871ms | ✓ 1639ms | http |
| 116.80.96.106:3172 | ✓ 1698ms | 否 | 否 | ✓ 1956ms | ✓ 1819ms | http |
| 103.113.70.189:1081 | ✓ 306ms | 否 | 否 | ✓ 992ms | ✓ 740ms | http |
| 61.52.131.172:8443 | ✓ 1002ms | ✓ 1349ms | ✓ 1074ms | ✓ 1371ms | ✓ 1070ms | http |
| 194.5.212.40:8080 | ✓ 1329ms | 否 | ✓ 867ms | ✓ 1964ms | ✓ 1620ms | http |

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
